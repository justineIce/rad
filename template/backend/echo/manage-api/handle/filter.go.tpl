package handle

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"{{.PackageName}}/utils"
	"{{.PackageName}}/utils/convert"
	"{{.PackageName}}/utils/enum"
	"{{.PackageName}}/manage-api/global"
	"{{.PackageName}}/model"
	"strings"
	"time"
)

const (
	//授权过期时间 2 小时
	AuthExpireTime int64 = 2*60*60 + 10
	//资源路径缓存10分钟
	PathAuthExpireTime int64 = 10 * 60
	//登录失效提示
	AuthInvalidMsg = "未登录或登陆已失效"
	//登录失败提示
	AuthLoginInfoErrorMsg = "获取登录信息失败"
	//登录cookie名称
	TokenName = "rad-d2_token"
	//资源缓存标记
	cacheResourceAllFlag = "cacheResourceAllFlag"
	//登录信息缓存标记
	loginInfoName = "rad-d2_loginInfo"
)

// Filter 路由拦截器
func Filter(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := GetToken(c)
		if err != nil {
			return utils.AuthFail(c, err.Error())
		}
		loginInfo, err := GetSysUserLoginInfo(token)
		if err != nil {
			return utils.AuthFail(c, err.Error())
		}
		var sysUser model.SysUser
		query := "id,company_id,office_id,username,name,photo,login_ip,login_date,login_flag,created_at,updated_at"
		if err := global.DB.Select(query).First(&sysUser, "id=?", loginInfo.ID).Error; err != nil {
			global.Log.Error("query sys_user error： %s", err.Error())
			return utils.AuthFail(c, AuthLoginInfoErrorMsg)
		}
		if sysUser.LoginFlag == enum.LoginFlag {
			return utils.AuthFail(c, "账号以被禁用")
		}
		//更新登录缓存时间
		sysUserFlag := GetSysUserLoginFlag(loginInfo.ID)
		_, err = global.RD.SetAndExpire(token, convert.MustJson(loginInfo), AuthExpireTime)
		if err != nil {
			global.Log.Error("redis SetAndExpire token error： %s", err.Error())
		}
		//单点登录统一标记
		_, err = global.RD.SetAndExpire(sysUserFlag, token, AuthExpireTime)
		if err != nil {
			global.Log.Error("redis SetAndExpire token error： %s", err.Error())
		}

		method := c.Request().Method
		path := c.Request().URL.Path
		//鉴定权限
		if !loginInfo.IsSuperAdmin {
			if v, ok := loginInfo.PathAuth[path+method]; ok && v.After(time.Now()) {
				//未过期
				global.Log.Info("鉴权未过期")
			} else {
				//获取所有鉴权列表，判断当前请求是否需要权限判定
				resourceAll, err := GetResourceAll()
				if err != nil {
					return utils.ErrorNull(c, "获取资源权限失败")
				}
				//非超级管理员，根据资源进行鉴权
				if checkResourceAllPath(resourceAll, method, path) && !checkResourcePath(loginInfo.SysMenuBtns, method, path) {
					return utils.PermissionDenied(c)
				}
				//权限缓存，10分钟
				loginInfo.PathAuth[path+method] = time.Now().Add(10 * time.Minute)
			}
		}
		c.Set(loginInfoName, &loginInfo)
		return next(c)
	}
}

// GetResourceAll 获取所有资源
func GetResourceAll() (btns []model.SysMenuBtn, err error) {
	cache, _ := global.RD.GetString(cacheResourceAllFlag)
	if cache == "" {
		if err = global.DB.Find(&btns).Error; err != nil {
			return
		}
		//redis缓存10分钟
		_, err = global.RD.SetAndExpire(cacheResourceAllFlag, convert.MustJsonString(btns), PathAuthExpireTime)
	} else {
		err = json.Unmarshal([]byte(cache), &btns)
	}
	return
}

// ClearResourceAll 清除资源
func ClearResourceAll() (err error) {
	_, err = global.RD.DelKey(cacheResourceAllFlag)
	if err != nil {
		global.Log.Error("redis DelKey CacheResourceAll error： %s", err.Error())
	}
	return
}

// checkResourceAllPath 判断请求的Url是否在鉴权范围内
func checkResourceAllPath(btns []model.SysMenuBtn, method string, path string) (flag bool) {
	for _, v := range btns {
		if strings.ToLower(v.Path) == strings.ToLower(path) && strings.ToLower(v.Method) == strings.ToLower(method) {
			//在鉴权范围内
			flag = true
		}
	}
	return
}

// checkResourcePath 判断用户的所有权限是否包含当前的请求Url
func checkResourcePath(btns []model.SysMenuBtn, method string, path string) (flag bool) {
	for _, v := range btns {
		if strings.ToLower(v.Path) == strings.ToLower(path) && strings.ToLower(v.Method) == strings.ToLower(method) {
			//在鉴权范围内
			flag = true
		}
	}
	return
}

// GetLoginInfo 登录信息
func GetLoginInfo(c echo.Context) *model.SysUserLoginInfo {
	loginInfo := c.Get(loginInfoName)
	if loginInfo == nil {
		return nil
	}
	return loginInfo.(*model.SysUserLoginInfo)
}

// GetSysUserLoginFlag 单点登录的唯一标记格式
func GetSysUserLoginFlag(id string) string {
	return fmt.Sprintf("sys_user_login_token_%s", id)
}

// GetToken 检查token有效性
func GetToken(c echo.Context) (tokenStr string, err error) {
	//auth
	tokenStr = c.Request().Header.Get("Access-Token")
	if tokenStr == "" {
		ck, err := c.Cookie(TokenName)
		if err == nil && ck.Value != "" {
			tokenStr = ck.Value
		}
	}
	if tokenStr == "" {
		tokenStr = c.FormValue(TokenName)
	}
	if tokenStr == "" {
		err = errors.New(AuthInvalidMsg)
		return
	}
	return
}

// GetSysUserLoginInfo 根据token获取Redis中的数据
//  tokenString：token参数
func GetSysUserLoginInfo(tokenStr string) (loginInfo model.SysUserLoginInfo, err error) {
	tip := AuthInvalidMsg
	loginInfoStr, err := global.RD.GetString(tokenStr)
	if err != nil {
		global.Log.Error("get token redis error： %s", err.Error())
		err = errors.New(AuthInvalidMsg)
	}
	if loginInfoStr == "" {
		err = errors.New(tip)
		return
	}
	err = json.Unmarshal([]byte(loginInfoStr), &loginInfo)
	return
}
