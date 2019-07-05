package handle

import (
	"github.com/LyricTian/gin-admin/pkg/errors"
	"github.com/jinzhu/gorm"
	"{{.PackageName}}/manage-api/global"
	"{{.PackageName}}/model"
	"strings"
)


// 根据创建人进行权限判断 - 数据操作权限判断【可以加缓存】
// userInfo 登录人
// createBy 数据的创建人
func PowerCheck(userInfo model.SysUserLoginInfo, createBy string) (err error) {
	var vSysUser model.VSysUser
	if err = global.DB.First(&vSysUser, "id=?", createBy).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.New("创建者不存在")
			return
		}
		global.Log.Error("PowerCheck error：", err)
		err = errors.New("获取创建者失败")
		return
	}
	if !userInfo.IsSuperAdmin && strings.Contains(vSysUser.OfficeRelationIds, userInfo.OfficeRelationIds) {
		err = errors.New("无权限操作")
		return
	}
	return
}
