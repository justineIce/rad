package handle

import (
	"{{.PackageName}}/manage-api/global"
	"{{.PackageName}}/model"
	"{{.PackageName}}/utils"
	"{{.PackageName}}/utils/validation"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// {{.TableRemark}}获取单条数据
func Get{{.StructName}}(c echo.Context) error {
	id := c.FormValue("id")
	var {{.SingName}} model.{{.StructName}}
	if err := global.DB.First(&{{.SingName}}, id).Error; err != nil {
		global.Log.Error("Get{{.StructName}} error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}
	{{ range .Columns }}{{if eq .ColumnName "created_by" }}
	loginInfo := GetLoginInfo(c)
	if err := PowerCheck(loginInfo, {{$.SingName}}.CreatedBy); err != nil {
		return utils.ErrorNull(c, err.Error())
	}{{end}}{{end}}
	return utils.SuccessNullMsg(c, {{.SingName}})
}

// {{.TableRemark}}获取所有数据
func Get{{.StructName}}All(c echo.Context) error {
	id := c.FormValue("id")
	var {{.SingName}} []model.{{.StructName}}
	if err := global.DB.Find(&{{.SingName}}, id).Error; err != nil {
		global.Log.Error("Get{{.StructName}}All error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}
	return utils.SuccessNullMsg(c, {{.SingName}})
}

// {{.TableRemark}}获取分页数据
func Get{{.StructName}}Page(c echo.Context) error {
    {{if not (eq .TableView "")}}

    db := global.DB.Model(&model.{{.ViewInfo.StructName}}{}).Order("id DESC")
    	{{if or .FieldsMap.office_id}}
        loginInfo := GetLoginInfo(c)
    	// 非超级管理员 && 数据权限验证
    	if !loginInfo.IsSuperAdmin {
            db = db.Select("{{.TableName}}.*").Joins("left join sys_office on sys_office.id=office_id").
            Where("relation_ids LIKE ?",loginInfo.OfficeRelationIds+"%")
    	}
    	{{end}}
    	/*	条件搜索范例
    		name := c.FormValue("name")
    		if name != "" {
    			db = db.Where("name like ?", fmt.Sprintf("%%s%%",name))
    		}
    	*/
    	var count int
    	var list []model.{{.ViewInfo.StructName}}
    	db.Count(&count)
    	pageIndex := utils.GetPageIndex(c.FormValue("page_index"))
    	pageSize := utils.GetPageSize(c.FormValue("page_size"))
    	if err := db.Limit(pageSize).Offset((pageIndex - 1) * pageSize).Scan(&list).Error; err != nil {
    		global.Log.Error("Get{{.ViewInfo.StructName}}Page error：%v", err)
    		return utils.ErrorNull(c, utils.GetFailResult)
    	}
    	return utils.SuccessNullMsg(c, &utils.PageData{
    		PageIndex:  pageIndex,
    		PageSize:   pageSize,
    		Count:      count,
    		PageNumber: utils.GetPageNumber(count, pageSize),
    		Data:       list,
    	})

    {{else}}

	db := global.DB.Model(&model.{{.StructName}}{}).Order("id DESC")
	{{if or .FieldsMap.office_id}}
    loginInfo := GetLoginInfo(c)
	// 非超级管理员 && 数据权限验证
	if !loginInfo.IsSuperAdmin {
        db = db.Select("{{.TableName}}.*").Joins("left join sys_office on sys_office.id=office_id").
        Where("relation_ids LIKE ?",loginInfo.OfficeRelationIds+"%")
	}
	{{end}}
	/*	条件搜索范例
		name := c.FormValue("name")
		if name != "" {
			db = db.Where("name like ?", fmt.Sprintf("%%s%%",name))
		}
	*/
	var count int
	var list []model.{{.StructName}}
	db.Count(&count)
	pageIndex := utils.GetPageIndex(c.FormValue("page_index"))
	pageSize := utils.GetPageSize(c.FormValue("page_size"))
	if err := db.Limit(pageSize).Offset((pageIndex - 1) * pageSize).Scan(&list).Error; err != nil {
		global.Log.Error("Get{{.StructName}}Page error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}
	return utils.SuccessNullMsg(c, &utils.PageData{
		PageIndex:  pageIndex,
		PageSize:   pageSize,
		Count:      count,
		PageNumber: utils.GetPageNumber(count, pageSize),
		Data:       list,
	})

	{{end}}
}

// {{.TableRemark}}保存数据
func Save{{.StructName}}(c echo.Context) error {
	{{if or .FieldsMap.updated_by .FieldsMap.created_by }}
	loginInfo := GetLoginInfo(c){{end}}
	var {{.SingName}}  model.{{.StructName}}
	idStr := c.FormValue("id")
	isEditFlag := false
	if idStr != "" {
		if err := global.DB.First(&{{.SingName}},  idStr).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}
        {{if .FieldsMap.created_by }}
        if err := PowerCheck(loginInfo, {{$.SingName}}.CreatedBy); err != nil {
            return utils.ErrorNull(c, err.Error())
        }{{end}}
		isEditFlag = true
	}else{
		//新增 {{.IDPrimaryKeyInt}}
		{{if .IDPrimaryKeyInt }}{{.SingName}}.ID = utils.ID(){{else}}{{.SingName}}.ID = utils.IDString(){{end}}
	}

    {{if .FieldsMap.updated_by }}
	{{$.SingName}}.UpdatedBy = loginInfo.ID{{end}}

	//需限制入参参数
	if err := new(utils.CustomBinder).Bind(&{{.SingName}}, c); err != nil {
		global.Log.Error(err.Error())
		return utils.ErrorNull(c,  utils.GetParsFailResult)
	}
	errs := validation.Valid(&{{.SingName}})
	if len(errs) > 0 {
		return utils.Error(c, "参数验证失败", errs)
	}
	if isEditFlag {
		//修改
		if err := global.DB.Save(&{{.SingName}}).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.UpdateFailResult)
		}
		return utils.SuccessNull(c, utils.UpdateSuccessResult)
	} else {
		if err := global.DB.Create(&{{.SingName}}).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c,   utils.AddFailResult)
		}
	    return utils.SuccessNull(c, utils.AddSuccessResult)
	}
}

// {{.TableRemark}}删除数据
func Del{{.StructName}}(c echo.Context) error {
	{{if or .FieldsMap.updated_by .FieldsMap.created_by }}
	loginInfo := GetLoginInfo(c){{end}}
	id := c.FormValue("id")
	if id == "" {
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	var {{.SingName}} model.{{.StructName}}
	if err := global.DB.Model(&model.{{.StructName}}{}).First(&{{.SingName}}, "id=?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}
		global.Log.Error("GetProjectType error：", err)
		return utils.ErrorNull(c, utils.GetDataFailResult)
	}
    {{if .FieldsMap.created_by }}
	if err := PowerCheck(loginInfo, {{$.SingName}}.CreatedBy); err != nil {
		return utils.ErrorNull(c, err.Error())
	}{{end}}

	if err := global.DB.Delete(&{{.SingName}}).Error; err != nil {
		global.Log.Error("Del{{.StructName}} error：%v", err)
		return utils.ErrorNull(c, utils.DeleteFailResult)
	}
	return utils.SuccessNull(c, utils.DeleteSuccessResult)
}
