package handle

import (
	"fmt"
	"{{.PackageName}}/manage-api/global"
	"{{.PackageName}}/model"
	"{{.PackageName}}/utils"
	"{{.PackageName}}/utils/validation"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// 机构表获取单条数据
func GetSysOffice(c echo.Context) error {
	id := c.FormValue("id")
	var sysOffice model.SysOffice
	if err := global.DB.First(&sysOffice, "id=?", id).Error; err != nil {
		global.Log.Error("GetSysOffice error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}

	loginInfo := GetLoginInfo(c)
	if err := PowerCheck(loginInfo, sysOffice.CreatedBy); err != nil {
		return utils.ErrorNull(c, err.Error())
	}
	return utils.SuccessNullMsg(c, sysOffice)
}

// 机构表获取所有数据
func GetSysOfficeAll(c echo.Context) error {
	loginInfo := GetLoginInfo(c)
	var sysOffice []model.SysOffice
	if err := global.DB.Order("sort ASC,created_at ASC").Find(&sysOffice, "sys_company_id=?", loginInfo.CompanyId).Error; err != nil {
		global.Log.Error("GetSysOfficeAll error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}
	return utils.SuccessNullMsg(c, model.GetSysOfficeTree(sysOffice, "0"))
}

// 机构表获取分页数据
func GetSysOfficePage(c echo.Context) error {
	db := global.DB.Model(&model.SysOffice{}).Order("id DESC")
	/*	条件搜索范例
		name := c.FormValue("name")
		if name != "" {
			db = db.Where("name like ?", fmt.Sprintf("%%s%%",name))
		}
	*/
	var count int
	var list []model.SysOffice
	db.Count(&count)
	pageIndex := utils.GetPageIndex(c.FormValue("page_index"))
	pageSize := utils.GetPageSize(c.FormValue("page_size"))
	if err := db.Limit(pageSize).Offset((pageIndex - 1) * pageSize).Scan(&list).Error; err != nil {
		global.Log.Error("GetSysOfficePage error：%v", err)
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

// 机构表保存数据
func SaveSysOffice(c echo.Context) error {

	loginInfo := GetLoginInfo(c)
	var sysOffice model.SysOffice
	idStr := c.FormValue("id")
	isEditFlag := false
	if idStr != "" {
		if err := global.DB.First(&sysOffice, "id=?", idStr).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}

		if err := PowerCheck(loginInfo, sysOffice.CreatedBy); err != nil {
			return utils.ErrorNull(c, err.Error())
		}
		isEditFlag = true
	} else {
		//新增 false
		sysOffice.ID = utils.IDString()
	}

	//需限制入参参数
	if err := new(utils.CustomBinder).Bind(&sysOffice, c); err != nil {
		global.Log.Error(err.Error())
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}

	sysOffice.UpdatedBy = loginInfo.ID
	if !isEditFlag {
		sysOffice.ID = utils.IDString()
		sysOffice.CreatedBy = loginInfo.ID
		sysOffice.SysCompanyID = loginInfo.CompanyId
	}
	if sysOffice.ParentID == "" {
		sysOffice.ParentID = "0"
	}
	if sysOffice.ParentID == "0" {
		sysOffice.RelationIds = fmt.Sprintf("-0-%s-", sysOffice.ID)
	} else {
		var parentSysOffice model.SysOffice
		if err := global.DB.Model(&model.SysMenu{}).First(&parentSysOffice, "id=?", sysOffice.ParentID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return utils.ErrorNull(c, utils.GetDataNullResult)
			}
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.GetDataFailResult)
		}
		sysOffice.RelationIds = fmt.Sprintf("%s%s-", parentSysOffice.RelationIds, sysOffice.ID)
	}

	errs := validation.Valid(&sysOffice)
	if len(errs) > 0 {
		return utils.Error(c, "参数验证失败", errs)
	}
	if isEditFlag {
		//修改
		if err := global.DB.Save(&sysOffice).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.UpdateFailResult)
		}
		return utils.SuccessNull(c, utils.UpdateSuccessResult)
	} else {
		if err := global.DB.Create(&sysOffice).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.AddFailResult)
		}
		return utils.SuccessNull(c, utils.AddSuccessResult)
	}
}

// 机构表删除数据
func DelSysOffice(c echo.Context) error {

	loginInfo := GetLoginInfo(c)
	id := c.FormValue("id")
	if id == "" {
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	var sysOffice model.SysOffice
	if err := global.DB.Model(&model.SysOffice{}).First(&sysOffice, "id=?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}
		global.Log.Error("GetProjectType error：", err)
		return utils.ErrorNull(c, utils.GetDataFailResult)
	}

	if err := PowerCheck(loginInfo, sysOffice.CreatedBy); err != nil {
		return utils.ErrorNull(c, err.Error())
	}

	//判断是否存在子部门
	var childrenCount int
	if err := global.DB.Model(&model.SysOffice{}).Where("parent_id=?", sysOffice.ID).Count(&childrenCount).Error; err != nil {
		global.Log.Error(err.Error())
		return utils.ErrorNull(c, utils.GetDataFailResult)
	}
	if childrenCount>0{
		return utils.ErrorNull(c, "旗下存在子部门无法进行删除！")
	}


	//判断是否存在用户
	if err := global.DB.Model(&model.SysUser{}).Where("office_id=?", sysOffice.ID).Count(&childrenCount).Error; err != nil {
		global.Log.Error(err.Error())
		return utils.ErrorNull(c, utils.GetDataFailResult)
	}
	if childrenCount>0{
		return utils.ErrorNull(c, "旗下存在用户无法进行删除！")
	}

	if err := global.DB.Delete(&sysOffice).Error; err != nil {
		global.Log.Error("DelSysOffice error：%v", err)
		return utils.ErrorNull(c, utils.DeleteFailResult)
	}
	return utils.SuccessNull(c, utils.DeleteSuccessResult)
}
