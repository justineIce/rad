package handle

import (
	"{{.PackageName}}/manage-api/global"
	"{{.PackageName}}/model"
	"{{.PackageName}}/utils"
	"{{.PackageName}}/utils/validation"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// 公司表获取单条数据
func GetSysCompany(c echo.Context) error {
	id := c.FormValue("id")
	var sysCompany model.SysCompany
	if err := global.DB.First(&sysCompany,"id=?", id).Error; err != nil {
		global.Log.Error("GetSysCompany error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}

	loginInfo := GetLoginInfo(c)
	if err := PowerCheck(loginInfo, sysCompany.CreatedBy); err != nil {
		return utils.ErrorNull(c, err.Error())
	}
	return utils.SuccessNullMsg(c, sysCompany)
}

// 公司表获取所有数据
func GetSysCompanyAll(c echo.Context) error {
	id := c.FormValue("id")
	var sysCompany []model.SysCompany
	if err := global.DB.Find(&sysCompany, id).Error; err != nil {
		global.Log.Error("GetSysCompanyAll error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}
	return utils.SuccessNullMsg(c, sysCompany)
}

// 公司表获取分页数据
func GetSysCompanyPage(c echo.Context) error {
	db := global.DB.Model(&model.SysCompany{}).Order("id DESC")
	/*	条件搜索范例
		name := c.FormValue("name")
		if name != "" {
			db = db.Where("name like ?", fmt.Sprintf("%%s%%",name))
		}
	*/
	var count int
	var list []model.SysCompany
	db.Count(&count)
	pageIndex := utils.GetPageIndex(c.FormValue("page_index"))
	pageSize := utils.GetPageSize(c.FormValue("page_size"))
	if err := db.Limit(pageSize).Offset((pageIndex - 1) * pageSize).Scan(&list).Error; err != nil {
		global.Log.Error("GetSysCompanyPage error：%v", err)
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

// 公司表保存数据
func SaveSysCompany(c echo.Context) error {

	loginInfo := GetLoginInfo(c)
	var sysCompany model.SysCompany
	idStr := c.FormValue("id")
	isEditFlag := false
	if idStr != "" {
		if err := global.DB.First(&sysCompany,"id=?", idStr).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}

		if err := PowerCheck(loginInfo, sysCompany.CreatedBy); err != nil {
			return utils.ErrorNull(c, err.Error())
		}
		isEditFlag = true
	} else {
		//新增 false
		sysCompany.ID = utils.IDString()
	}

	sysCompany.UpdatedBy = loginInfo.ID

	//需限制入参参数
	if err := new(utils.CustomBinder).Bind(&sysCompany, c); err != nil {
		global.Log.Error(err.Error())
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	errs := validation.Valid(&sysCompany)
	if len(errs) > 0 {
		return utils.Error(c, "参数验证失败", errs)
	}
	if isEditFlag {
		//修改
		if err := global.DB.Save(&sysCompany).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.UpdateFailResult)
		}
		return utils.SuccessNull(c, utils.UpdateSuccessResult)
	} else {
		if err := global.DB.Create(&sysCompany).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.AddFailResult)
		}
		return utils.SuccessNull(c, utils.AddSuccessResult)
	}
}

// 公司表删除数据
func DelSysCompany(c echo.Context) error {

	loginInfo := GetLoginInfo(c)
	id := c.FormValue("id")
	if id == "" {
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	var sysCompany model.SysCompany
	if err := global.DB.Model(&model.SysCompany{}).First(&sysCompany, "id=?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}
		global.Log.Error("GetProjectType error：", err)
		return utils.ErrorNull(c, utils.GetDataFailResult)
	}

	if err := PowerCheck(loginInfo, sysCompany.CreatedBy); err != nil {
		return utils.ErrorNull(c, err.Error())
	}

	if err := global.DB.Delete(&sysCompany).Error; err != nil {
		global.Log.Error("DelSysCompany error：%v", err)
		return utils.ErrorNull(c, utils.DeleteFailResult)
	}
	return utils.SuccessNull(c, utils.DeleteSuccessResult)
}
