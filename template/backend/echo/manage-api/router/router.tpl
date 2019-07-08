package router

import (
	"github.com/labstack/echo"
	"{{.PackageName}}/manage-api/handle"
)

func {{.StructName}}Router(e *echo.Group) {
	e.POST("/{{.TableName}}/get", handle.Get{{.StructName}})
	e.POST("/{{.TableName}}/all", handle.Get{{.StructName}}All)
	e.POST("/{{.TableName}}/page", handle.Get{{.StructName}}Page)
	e.POST("/{{.TableName}}/save", handle.Save{{.StructName}})
	e.POST("/{{.TableName}}/del", handle.Del{{.StructName}})
}
