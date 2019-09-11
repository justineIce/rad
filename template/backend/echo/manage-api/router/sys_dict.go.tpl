package router

import (
	"{{.PackageName}}/manage-api/handle"
	"github.com/labstack/echo"
)

func SysDictRouter(e *echo.Group) {
	e.POST("/sys_dict/get", handle.GetSysDict)
	e.POST("/sys_dict/all", handle.GetSysDictAll)
	e.POST("/sys_dict/page", handle.GetSysDictPage)
	e.POST("/sys_dict/save", handle.SaveSysDict)
	e.POST("/sys_dict/del", handle.DelSysDict)
}
