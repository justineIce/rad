package handle

import (
	"{{.PackageName}}/manage-api/global"
	"{{.PackageName}}/model"
	"{{.PackageName}}/utils"
	"{{.PackageName}}/utils/validation"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// 角色表获取单条数据
func GetSysRole(c echo.Context) error {
	id := c.FormValue("id")
	var sysRole model.SysRole
	if err := global.DB.First(&sysRole,"id=?", id).Error; err != nil {
		global.Log.Error("GetSysRole error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}

	loginInfo := GetLoginInfo(c)
	if err := PowerCheck(loginInfo, sysRole.CreatedBy); err != nil {
		return utils.ErrorNull(c, err.Error())
	}
	return utils.SuccessNullMsg(c, sysRole)
}

// 角色表获取所有数据
func GetSysRoleAll(c echo.Context) error {
	id := c.FormValue("id")
	var sysRole []model.SysRole
	if err := global.DB.Find(&sysRole, id).Error; err != nil {
		global.Log.Error("GetSysRoleAll error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}
	return utils.SuccessNullMsg(c, sysRole)
}

// 角色表获取分页数据
func GetSysRolePage(c echo.Context) error {
	db := global.DB.Model(&model.SysRole{}).Order("id DESC")
	/*	条件搜索范例
		name := c.FormValue("name")
		if name != "" {
			db = db.Where("name like ?", fmt.Sprintf("%%s%%",name))
		}
	*/
	var count int
	var list []model.SysRole
	db.Count(&count)
	pageIndex := utils.GetPageIndex(c.FormValue("page_index"))
	pageSize := utils.GetPageSize(c.FormValue("page_size"))
	if err := db.Limit(pageSize).Offset((pageIndex - 1) * pageSize).Scan(&list).Error; err != nil {
		global.Log.Error("GetSysRolePage error：%v", err)
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

// 角色表保存数据
func SaveSysRole(c echo.Context) error {

	loginInfo := GetLoginInfo(c)
	var sysRole model.SysRole
	idStr := c.FormValue("id")
	isEditFlag := false
	if idStr != "" {
		if err := global.DB.First(&sysRole,"id=?", idStr).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}

		if err := PowerCheck(loginInfo, sysRole.CreatedBy); err != nil {
			return utils.ErrorNull(c, err.Error())
		}
		isEditFlag = true
	} else {
		//新增 false
		sysRole.ID = utils.IDString()
	}

	sysRole.UpdatedBy = loginInfo.ID

	//需限制入参参数
	if err := new(utils.CustomBinder).Bind(&sysRole, c); err != nil {
		global.Log.Error(err.Error())
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	if sysRole.RoleType == model.UserPostSuperAdminCode {
		return utils.ErrorNull(c, "无权限添加此角色类型")
	}
	//业务逻辑
	sysRole.UpdatedBy = loginInfo.ID
	if !isEditFlag {
		sysRole.ID = utils.IDString()
		sysRole.CreatedBy = loginInfo.ID
		sysRole.SysOfficeID = loginInfo.OfficeId
	}

	errs := validation.Valid(&sysRole)
	if len(errs) > 0 {
		return utils.Error(c, "参数验证失败", errs)
	}
	if isEditFlag {
		//修改
		if err := global.DB.Save(&sysRole).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.UpdateFailResult)
		}
		return utils.SuccessNull(c, utils.UpdateSuccessResult)
	} else {
		if err := global.DB.Create(&sysRole).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.AddFailResult)
		}
		return utils.SuccessNull(c, utils.AddSuccessResult)
	}
}

// 角色表删除数据
func DelSysRole(c echo.Context) error {

	loginInfo := GetLoginInfo(c)
	id := c.FormValue("id")
	if id == "" {
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	var sysRole model.SysRole
	if err := global.DB.Model(&model.SysRole{}).First(&sysRole, "id=?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}
		global.Log.Error("GetProjectType error：", err)
		return utils.ErrorNull(c, utils.GetDataFailResult)
	}

	if err := PowerCheck(loginInfo, sysRole.CreatedBy); err != nil {
		return utils.ErrorNull(c, err.Error())
	}

	if err := global.DB.Delete(&sysRole).Error; err != nil {
		global.Log.Error("DelSysRole error：%v", err)
		return utils.ErrorNull(c, utils.DeleteFailResult)
	}
	return utils.SuccessNull(c, utils.DeleteSuccessResult)
}
