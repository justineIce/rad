package handle

import (
	"github.com/labstack/echo"
	"{{.PackageName}}/global"
	"{{.PackageName}}/model"
	"{{.PackageName}}/utils"
)

// {{.TableRemark}}获取单条数据
// @Param {paramName} query string false "{paramDesc}"
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
// @Param {paramName} query string false "{paramDesc}"
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
// @Param {paramName} query string false "{paramDesc}"
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
	pageIndex := model.GetPageIndex(c.FormValue("page_index"))
	pageSize := model.GetPageSize(c.FormValue("page_size"))
	if err := db.Limit(pageSize).Offset((pageIndex - 1) * pageSize).Scan(&list).Error; err != nil {
		c.Logger().Errorf("Get{{.StructName}}Page error：%v", err)
		return utils.ErrorNull(c, utils.GetFailResult)
	}
	return utils.SuccessNullMsg(c, &model.PageData{
		PageIndex:  pageIndex,
		PageSize:   pageSize,
		Count:      count,
		PageNumber: model.GetPageNumber(count, pageSize),
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
			c.Logger().Error(err)
			return utils.ErrorNull(c, utils.GetDataNullResult)
		}
		isFlag = true
	}
	//需限制入参参数
	if err := new(CustomBinder).Bind(&{{.singName}}, c); err != nil {
		c.Logger().Error(err)
		return utils.ErrorNull(c,  utils.GetParsFailResult)
	}
	if isFlag {
		//修改
		if err := global.DB.Save(&{{.singName}}).Error; err != nil {
			c.Logger().Error(err)
			return utils.ErrorNull(c, utils.UpdateFailResult)
		}
		return utils.SuccessNullMsg(c, utils.UpdateSuccessResult)
	} else {
		//新增
		{{.singName}}.ID = utils.ID()
		if err := global.DB.Create(&{{.singName}}).Error; err != nil {
			c.Logger().Error(err)
			return utils.ErrorNull(c,   utils.AddFailResult)
		}
	    return utils.SuccessNullMsg(c, utils.AddSuccessResult)
	}
}

// {{.TableRemark}}新增数据
func Add{{.StructName}}(c echo.Context) error {
	{{.singName}} := new(model.{{.StructName}})
	if err := c.Bind({{.singName}}); err != nil {
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}

	//判断是否自增
	{{.singName}}.ID = utils.ID()

	if err := global.DB.Create({{.singName}}).Error; err != nil {
		c.Logger().Errorf("Add{{.StructName}} error：%v", err)
		return utils.ErrorNull(c, utils.AddFailResult)
	}
	return utils.SuccessNullMsg(c, utils.AddSuccessResult)
}

// {{.TableRemark}}修改数据
func Update{{.StructName}}(c echo.Context) error {
	{{.singName}} := new(model.{{.StructName}})
	if err := c.Bind({{.singName}}); err != nil {
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	if err := global.DB.Save({{.singName}}).Error; err != nil {
		c.Logger().Errorf("Update{{.StructName}} error：%v", err)
		return utils.ErrorNull(c, utils.UpdateFailResult)
	}
	return utils.SuccessNullMsg(c, utils.UpdateSuccessResult)
}

// {{.TableRemark}}删除数据
func Del{{.StructName}}(c echo.Context) error {
	{{.singName}} := new(model.{{.StructName}})
	if err := c.Bind({{.singName}}); err != nil {
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	if {{.singName}}.ID < 0 {
		return utils.ErrorNull(c, utils.DeleteFailResult)
	}
	if err := global.DB.Delete({{.singName}}).Error; err != nil {
		c.Logger().Errorf("Del{{.StructName}} error：%v", err)
		return utils.ErrorNull(c, utils.DeleteFailResult)
	}
	return utils.SuccessNullMsg(c, utils.DeleteSuccessResult)
}
