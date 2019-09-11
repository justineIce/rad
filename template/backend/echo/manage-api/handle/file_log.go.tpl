package handle

import (
	"{{.PackageName}}/manage-api/global"
	"{{.PackageName}}/model"
	"{{.PackageName}}/utils"
	"{{.PackageName}}/utils/validation"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// 获取单条数据
func GetFileLog(c echo.Context) error {
	id := c.FormValue("id")
	var fileLog model.FileLog
	if err := global.DB.First(&fileLog,"id=?", id).Error; err != nil {
		global.Log.Error("GetFileLog error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}

	loginInfo := GetLoginInfo(c)
	if err := PowerCheck(loginInfo, fileLog.CreatedBy); err != nil {
		return utils.ErrorNull(c, err.Error())
	}
	return utils.SuccessNullMsg(c, fileLog)
}

// 获取所有数据
func GetFileLogAll(c echo.Context) error {
	id := c.FormValue("id")
	var fileLog []model.FileLog
	if err := global.DB.Find(&fileLog, id).Error; err != nil {
		global.Log.Error("GetFileLogAll error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}
	return utils.SuccessNullMsg(c, fileLog)
}

// 获取分页数据
func GetFileLogPage(c echo.Context) error {
	db := global.DB.Model(&model.FileLog{}).Order("id DESC")
	/*	条件搜索范例
		name := c.FormValue("name")
		if name != "" {
			db = db.Where("name like ?", fmt.Sprintf("%%s%%",name))
		}
	*/
	var count int
	var list []model.FileLog
	db.Count(&count)
	pageIndex := utils.GetPageIndex(c.FormValue("page_index"))
	pageSize := utils.GetPageSize(c.FormValue("page_size"))
	if err := db.Limit(pageSize).Offset((pageIndex - 1) * pageSize).Scan(&list).Error; err != nil {
		global.Log.Error("GetFileLogPage error：%v", err)
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

// 保存数据
func SaveFileLog(c echo.Context) error {

	loginInfo := GetLoginInfo(c)
	var fileLog model.FileLog
	idStr := c.FormValue("id")
	isEditFlag := false
	if idStr != "" {
		if err := global.DB.First(&fileLog,"id=?", idStr).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}

		if err := PowerCheck(loginInfo, fileLog.CreatedBy); err != nil {
			return utils.ErrorNull(c, err.Error())
		}
		isEditFlag = true
	} else {
		//新增 false
		fileLog.ID = utils.IDString()
	}

	//需限制入参参数
	if err := new(utils.CustomBinder).Bind(&fileLog, c); err != nil {
		global.Log.Error(err.Error())
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}

	fileLog.UpdatedBy = loginInfo.ID
	if !isEditFlag {
		fileLog.ID = utils.IDString()
		fileLog.CreatedBy = loginInfo.ID
	}

	errs := validation.Valid(&fileLog)
	if len(errs) > 0 {
		return utils.Error(c, "参数验证失败", errs)
	}
	if isEditFlag {
		//修改
		if err := global.DB.Save(&fileLog).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.UpdateFailResult)
		}
		return utils.SuccessNull(c, utils.UpdateSuccessResult)
	} else {
		if err := global.DB.Create(&fileLog).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.AddFailResult)
		}
		return utils.SuccessNull(c, utils.AddSuccessResult)
	}
}

// 删除数据
func DelFileLog(c echo.Context) error {

	loginInfo := GetLoginInfo(c)
	id := c.FormValue("id")
	if id == "" {
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	var fileLog model.FileLog
	if err := global.DB.Model(&model.FileLog{}).First(&fileLog, "id=?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}
		global.Log.Error("GetProjectType error：", err)
		return utils.ErrorNull(c, utils.GetDataFailResult)
	}

	if err := PowerCheck(loginInfo, fileLog.CreatedBy); err != nil {
		return utils.ErrorNull(c, err.Error())
	}

	if err := global.DB.Delete(&fileLog).Error; err != nil {
		global.Log.Error("DelFileLog error：%v", err)
		return utils.ErrorNull(c, utils.DeleteFailResult)
	}
	return utils.SuccessNull(c, utils.DeleteSuccessResult)
}
