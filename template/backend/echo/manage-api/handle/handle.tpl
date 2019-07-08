package handle

import (
	"{{.PackageName}}/manage-api/global"
	"{{.PackageName}}/model"
	"{{.PackageName}}/utils"
	"{{.PackageName}}/utils/validation"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"time"
)

// {{.TableRemark}}获取单条数据
func Get{{.StructName}}(c echo.Context) error {
	id := c.FormValue("id")
	var {{.singName}} model.{{.StructName}}
	if err := global.DB.First(&{{.singName}}, id).Error; err != nil {
		c.Logger().Errorf("Get{{.StructName}} error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}
	{{ range .Columns }}{{if eq .ColumnName "created_by" }}
	loginInfo := GetLoginInfo(c)
	if err := PowerCheck(loginInfo, {{$.singName}}.CreatedBy); err != nil {
		return utils.ErrorNull(c, err.Error())
	}{{end}}{{end}}
	return utils.SuccessNullMsg(c, {{.singName}})
}

// {{.TableRemark}}获取所有数据
func Get{{.StructName}}All(c echo.Context) error {
	id := c.FormValue("id")
	var {{.singName}} []model.{{.StructName}}
	if err := global.DB.Find(&{{.singName}}, id).Error; err != nil {
		c.Logger().Errorf("Get{{.StructName}}All error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}
	return utils.SuccessNullMsg(c, {{.singName}})
}

// {{.TableRemark}}获取分页数据
func Get{{.StructName}}Page(c echo.Context) error {
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
		c.Logger().Errorf("Get{{.StructName}}Page error：%v", err)
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

// {{.TableRemark}}保存数据
func Save{{.StructName}}(c echo.Context) error {
	{{if or .FieldsMap.updated_by .FieldsMap.created_by }}
	loginInfo := GetLoginInfo(c){{end}}
	var {{.singName}}  model.{{.StructName}}
	idStr := c.FormValue("id")
	isEditFlag := false
	if idStr != "" {
		if err := global.DB.First(&{{.singName}},  idStr).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}
        {{if .FieldsMap.created_by }}
        if err := PowerCheck(loginInfo, {{$.singName}}.CreatedBy); err != nil {
            return utils.ErrorNull(c, err.Error())
        }{{end}}
		isEditFlag = true
	}else{
		//新增 {{.IDPrimaryKeyInt}}
		{{if .IDPrimaryKeyInt }}{{.singName}}.ID = utils.ID(){{else}}{{.singName}}.ID = utils.IDString(){{end}}
	}

    {{if .FieldsMap.updated_by }}
	{{$.singName}}.UpdatedBy = loginInfo.ID{{end}}

	//需限制入参参数
	if err := new(utils.CustomBinder).Bind(&{{.singName}}, c); err != nil {
		global.Log.Error(err.Error())
		return utils.ErrorNull(c,  utils.GetParsFailResult)
	}
	errs := validation.Valid(&{{.singName}})
	if len(errs) > 0 {
		return utils.Error(c, "参数验证失败", errs)
	}
	if isEditFlag {
		//修改
		if err := global.DB.Save(&{{.singName}}).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.UpdateFailResult)
		}
		return utils.SuccessNull(c, utils.UpdateSuccessResult)
	} else {
		if err := global.DB.Create(&{{.singName}}).Error; err != nil {
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
	var {{.singName}} model.{{.StructName}}
	if err := global.DB.Model(&model.{{.StructName}}{}).First(&{{.singName}}, "id=?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}
		global.Log.Error("GetProjectType error：", err)
		return utils.ErrorNull(c, utils.GetDataFailResult)
	}
    {{if .FieldsMap.created_by }}
	if err := PowerCheck(loginInfo, {{$.singName}}.CreatedBy); err != nil {
		return utils.ErrorNull(c, err.Error())
	}{{end}}

    {{if .FieldsMap.updated_by }}
	if err := global.DB.Model(&{{.singName}}).Updates(map[string]interface{}{"updated_by": loginInfo.ID, "deleted_at": utils.FormatTime(time.Now())}); err != nil {
		c.Logger().Errorf("Del{{.StructName}} error：%v", err)
		return utils.ErrorNull(c, utils.DeleteFailResult)
	}
	{{else}}
	if err := global.DB.Model(&{{.singName}}).Updates(map[string]interface{}{"deleted_at": utils.FormatTime(time.Now())}); err != nil {
		c.Logger().Errorf("Del{{.StructName}} error：%v", err)
		return utils.ErrorNull(c, utils.DeleteFailResult)
	}
	{{end}}
	return utils.SuccessNull(c, utils.DeleteSuccessResult)
}
