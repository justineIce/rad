package handle

import (
	"{{.PackageName}}/manage-api/global"
	"{{.PackageName}}/model"
	"{{.PackageName}}/utils"
	"{{.PackageName}}/utils/validation"
	"github.com/labstack/echo"
)

// {{.TableRemark}}获取单条数据
func Get{{.StructName}}(c echo.Context) error {
	id := c.FormValue("id")
	var {{.singName}} model.{{.StructName}}
	if err := global.DB.First(&{{.singName}}, id).Error; err != nil {
		c.Logger().Errorf("Get{{.StructName}} error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}
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
	var {{.singName}}  model.{{.StructName}}
	idStr := c.FormValue("id")
	isFlag := false
	if idStr != "" {
		if err := global.DB.First(&{{.singName}},  idStr).Error; err != nil {
			global.Log.Error(err.Error())
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}
		isFlag = true
	}else{
		//新增 {{.IDPrimaryKeyInt}}
		{{if .IDPrimaryKeyInt }}{{.singName}}.ID = utils.ID(){{else}}{{.singName}}.ID = utils.IDString(){{end}}
	}
	//需限制入参参数
	if err := new(utils.CustomBinder).Bind(&{{.singName}}, c); err != nil {
		global.Log.Error(err.Error())
		return utils.ErrorNull(c,  utils.GetParsFailResult)
	}
	errs := validation.Valid(&{{.singName}})
	if len(errs) > 0 {
		return utils.Error(c, "参数验证失败", errs)
	}
	if isFlag {
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
	{{.singName}} := new(model.{{.StructName}})
	if err := c.Bind({{.singName}}); err != nil {
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	{{if .IDPrimaryKeyInt }}
	if {{.singName}}.ID < 0 {
    	return utils.ErrorNull(c, utils.DeleteFailResult)
    }
    {{else}}
    if {{.singName}}.ID == "" {
        return utils.ErrorNull(c, utils.DeleteFailResult)
    }
    {{end}}
	if err := global.DB.Delete({{.singName}}).Error; err != nil {
		c.Logger().Errorf("Del{{.StructName}} error：%v", err)
		return utils.ErrorNull(c, utils.DeleteFailResult)
	}
	return utils.SuccessNullMsg(c, utils.DeleteSuccessResult)
}
