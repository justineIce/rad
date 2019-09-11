package handle

import (
	"fmt"
	"{{.PackageName}}/manage-api/global"
	"{{.PackageName}}/model"
	"{{.PackageName}}/utils"
	"{{.PackageName}}/utils/validation"
	"github.com/labstack/echo"
	"strings"
)

// 日志表获取分页数据
func GetSysLogPage(c echo.Context) error {
	loginInfo := GetLoginInfo(c)
	//排序
	order := "created_at DESC"
	sortField := c.FormValue("sortField")
	sortOrder := c.FormValue("sortOrder")
	if sortField != "" && sortOrder != "" {
		order = fmt.Sprintf("%s %s", sortField, sortOrder)
	}
	db := global.DB.Model(&model.VSysLog{}).Order(order)
	if !loginInfo.IsSuperAdmin {
		db = db.Where("company_id=?", loginInfo.CompanyId)
	}
	//条件搜索范例
	name := c.FormValue("name")
	createdBy := c.FormValue("created_by")
	t := c.FormValue("type")
	url := c.FormValue("url")
	ip := c.FormValue("ip")
	title := c.FormValue("title")
	startTime := c.FormValue("time1")
	endTime := c.FormValue("time2")
	if name != "" {
		db = db.Where("name like ?", fmt.Sprintf("%%%s%%", name))
	}
	if title != "" {
		db = db.Where("title like ?", fmt.Sprintf("%%%s%%", title))
	}
	if url != "" {
		db = db.Where("request_url like ?", fmt.Sprintf("%%%s%%", url))
	}
	if ip != "" {
		db = db.Where("remote_addr like ?", fmt.Sprintf("%%%s%%", ip))
	}
	if t != "" {
		db = db.Where("type = ?", t)
	}
	if createdBy != "" {
		db = db.Where("created_by=?", createdBy)
	}
	if startTime != "" && endTime != "" {
		db = db.Where("created_at>=? AND created_at<=?", startTime, endTime)
	}
	var count int
	var list []model.VSysLog
	var err error
	db.Count(&count)
	pageIndex := utils.GetPageIndex(c.FormValue("pageNo"))
	pageSize := utils.GetPageSize(c.FormValue("pageSize"))
	if err = db.Limit(pageSize).Offset((pageIndex - 1) * pageSize).Scan(&list).Error; err != nil {
		global.Log.Error("GetSysLogPage error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}
	return utils.SuccessNullMsg(c, &utils.PageData{
		PageIndex:  pageIndex,
		PageSize:   pageSize,
		Count:      count,
		PageNumber: utils.GetPageNumber(count, pageSize),
		Data:       list,
	})

}

var LogMap = map[string]string{
	"/api/auth/file_log/del":           "删除上传文件日志",
	"/api/auth/file_log/save":          "保存上传文件日志",
	"/api/auth/files/remove":           "删除文件",
	"/api/auth/files/upload/:auth":     "删除文件",
	"/api/auth/sys_company/del":        "删除公司信息",
	"/api/auth/sys_company/save":       "保存公司信息",
	"/api/auth/sys_dict/del":           "删除数据字典",
	"/api/auth/sys_dict/save":          "保存数据字典",
	"/api/auth/sys_menu/del":           "删除系统菜单",
	"/api/auth/sys_menu/save":          "保存系统菜单",
	"/api/auth/sys_menu_btn/del":       "删除菜单按钮",
	"/api/auth/sys_menu_btn/save":      "保存菜单按钮",
	"/api/auth/sys_office/del":         "删除组织架构",
	"/api/auth/sys_office/save":        "保存组织架构",
	"/api/auth/sys_role/del":           "删除角色",
	"/api/auth/sys_role/save":          "保存角色",
	"/api/auth/sys_role_menu/del":      "删除角色菜单",
	"/api/auth/sys_role_menu/save":     "保存角色菜单",
	"/api/auth/sys_role_menu_btn/del":  "删除角色菜单按钮",
	"/api/auth/sys_role_menu_btn/save": "保存角色菜单按钮",
	"/api/auth/sys_user/del":           "删除系统用户",
	"/api/auth/sys_user/pwd/reset":     "重置系统用户密码",
	"/api/auth/sys_user/save":          "保存系统用户",
	"/api/auth/sys_user_role/del":      "删除系统用户角色",
	"/api/auth/sys_user_role/save":     "保存系统用户角色",
	"/api/login":                       "系统账号登录",
	"/api/logout":                      "系统账号登出",
}

// 日志表保存数据
func SaveSysLog(c echo.Context, result utils.ResultParam) (err error) {
	defer func() {
		if errRec := recover(); err != nil {
			err = fmt.Errorf("异常：%v", errRec)
			global.Log.Error(err.Error())
		}
	}()
	path := c.Request().URL.Path
	title, ok := LogMap[path]
	if ok {

		var loginInfoID string
		var companyID string
		loginInfo := GetLoginInfo(c)
		if loginInfo == nil {
			loginID := c.Get("loginID")
			if loginID != nil {
				loginInfoID = loginID.(string)
			}
			companyIDObj := c.Get("CompanyID")
			if companyIDObj != nil {
				companyID = companyIDObj.(string)
			}
		} else {
			loginInfoID = loginInfo.ID
			companyID = loginInfo.CompanyId
		}
		if loginInfoID == "" || companyID == "" {
			return
		}

		var sysLog model.SysLog
		//新增 false
		sysLog.ID = utils.IDString()
		sysLog.Title = title
		sysLog.UpdatedBy = loginInfoID
		sysLog.ID = utils.IDString()
		sysLog.CreatedBy = loginInfoID
		sysLog.RemoteAddr = c.RealIP()
		sysLog.UserAgent = c.Request().UserAgent()
		sysLog.RequestURL = c.Path()
		sysLog.Method = c.Request().Method
		sysLog.ResultStatus = result.Ret
		sysLog.ResultMsg = result.Msg
		sysLog.SysCompanyID = companyID
		uv, _ := c.FormParams()
		var arr []string
		for k, v := range uv {
			if k == "password" {
				arr = append(arr, "******")
			} else {
				arr = append(arr, fmt.Sprintf("%s=%v", k, strings.Join(v, "")))
			}
		}
		sysLog.Params = strings.Join(arr, "&")
		errs := validation.Valid(&sysLog)
		if len(errs) > 0 {
			err = fmt.Errorf("参数验证失败，%v", errs)
			global.Log.Error(err.Error())
			return
		}
		if err = global.DB.Create(&sysLog).Error; err != nil {
			global.Log.Error(err.Error())
		}
	}
	return
}
