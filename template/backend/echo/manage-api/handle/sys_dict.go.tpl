package handle

import (
	"{{.PackageName}}/manage-api/global"
	"{{.PackageName}}/model"
	"{{.PackageName}}/utils"
	"{{.PackageName}}/utils/validation"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// 字典表获取单条数据
func GetSysDict(c echo.Context) error {
	id := c.FormValue("id")
	var sysDict model.SysDict
	if err := global.DB.First(&sysDict,"id=?", id).Error; err != nil {
		global.Log.Error("GetSysDict error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}

	loginInfo := GetLoginInfo(c)
	if err := PowerCheck(loginInfo, sysDict.CreatedBy); err != nil {
		return utils.ErrorNull(c, err.Error())
	}
	return utils.SuccessNullMsg(c, sysDict)
}

// 字典表获取所有数据
func GetSysDictAll(c echo.Context) error {
	id := c.FormValue("id")
	var sysDict []model.SysDict
	if err := global.DB.Find(&sysDict, id).Error; err != nil {
		global.Log.Error("GetSysDictAll error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}
	return utils.SuccessNullMsg(c, sysDict)
}

// 字典表获取分页数据
func GetSysDictPage(c echo.Context) error {
	db := global.DB.Model(&model.SysDict{}).Order("id DESC")
	/*	条件搜索范例
		name := c.FormValue("name")
		if name != "" {
			db = db.Where("name like ?", fmt.Sprintf("%%s%%",name))
		}
	*/
	var count int
	var list []model.SysDict
	db.Count(&count)
	pageIndex := utils.GetPageIndex(c.FormValue("page_index"))
	pageSize := utils.GetPageSize(c.FormValue("page_size"))
	if err := db.Limit(pageSize).Offset((pageIndex - 1) * pageSize).Scan(&list).Error; err != nil {
		global.Log.Error("GetSysDictPage error：%v", err)
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

// 字典表保存数据
func SaveSysDict(c echo.Context) error {
	loginInfo := GetLoginInfo(c)
	var sysDict model.SysDict
	idStr := c.FormValue("id")
	isEditFlag := false
	if idStr != "" {
		if err := global.DB.First(&sysDict,"id=?", idStr).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}

		if err := PowerCheck(loginInfo, sysDict.CreatedBy); err != nil {
			return utils.ErrorNull(c, err.Error())
		}
		isEditFlag = true
	} else {
		//新增 false
		sysDict.ID = utils.IDString()
	}

	sysDict.UpdatedBy = loginInfo.ID

	//需限制入参参数
	if err := new(utils.CustomBinder).Bind(&sysDict, c); err != nil {
		global.Log.Error(err.Error())
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	errs := validation.Valid(&sysDict)
	if len(errs) > 0 {
		return utils.Error(c, "参数验证失败", errs)
	}
	if isEditFlag {
		//修改
		if err := global.DB.Save(&sysDict).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.UpdateFailResult)
		}
		return utils.SuccessNull(c, utils.UpdateSuccessResult)
	} else {
		if err := global.DB.Create(&sysDict).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.AddFailResult)
		}
		return utils.SuccessNull(c, utils.AddSuccessResult)
	}
}

// 字典表删除数据
func DelSysDict(c echo.Context) error {

	loginInfo := GetLoginInfo(c)
	id := c.FormValue("id")
	if id == "" {
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	var sysDict model.SysDict
	if err := global.DB.Model(&model.SysDict{}).First(&sysDict, "id=?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}
		global.Log.Error("GetProjectType error：", err)
		return utils.ErrorNull(c, utils.GetDataFailResult)
	}

	if err := PowerCheck(loginInfo, sysDict.CreatedBy); err != nil {
		return utils.ErrorNull(c, err.Error())
	}

	if err := global.DB.Delete(&sysDict).Error; err != nil {
		global.Log.Error("DelSysDict error：%v", err)
		return utils.ErrorNull(c, utils.DeleteFailResult)
	}
	return utils.SuccessNull(c, utils.DeleteSuccessResult)
}
