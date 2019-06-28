package router

import (
	"github.com/labstack/echo"
	"{{.PackageName}}/handle"
)

func {{.StructName}}Router(e *echo.Group) {
	e.POST("/{{.SingName}}/get", handle.Get{{.StructName}})
	e.POST("/{{.SingName}}/all", handle.Get{{.StructName}}All)
	e.POST("/{{.SingName}}/page", handle.Get{{.StructName}}Page)
	e.POST("/{{.SingName}}/add", handle.Add{{.StructName}})
	e.POST("/{{.SingName}}/update", handle.Update{{.StructName}})
	e.POST("/{{.SingName}}/del", handle.Del{{.StructName}})
}
