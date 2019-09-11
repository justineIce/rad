package router

import (
	"{{.PackageName}}/manage-api/handle"
	"github.com/labstack/echo"
)

func SysMenuRouter(e *echo.Group) {
	e.POST("/sys_menu/get", handle.GetSysMenu)
	e.POST("/sys_menu/all", handle.GetSysMenuAll)
	e.POST("/sys_menu/page", handle.GetSysMenuPage)
	e.POST("/sys_menu/save", handle.SaveSysMenu)
	e.POST("/sys_menu/del", handle.DelSysMenu)
}
