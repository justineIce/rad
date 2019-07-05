package router

import (
	"github.com/labstack/echo"
	"{{.PackageName}}/manage-api/handle"
)

func {{.StructName}}Router(e *echo.Group) {
	e.POST("/{{.SingName}}/get", handle.Get{{.StructName}})
	e.POST("/{{.SingName}}/all", handle.Get{{.StructName}}All)
	e.POST("/{{.SingName}}/page", handle.Get{{.StructName}}Page)
	e.POST("/{{.SingName}}/save", handle.Save{{.StructName}})
	e.POST("/{{.SingName}}/del", handle.Del{{.StructName}})
}
