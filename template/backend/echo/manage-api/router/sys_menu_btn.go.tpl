package router

import (
	"{{.PackageName}}/manage-api/handle"
	"github.com/labstack/echo"
)

func SysMenuBtnRouter(e *echo.Group) {
	e.POST("/sys_menu_btn/get", handle.GetSysMenuBtn)
	e.POST("/sys_menu_btn/all", handle.GetSysMenuBtnAll)
	e.POST("/sys_menu_btn/page", handle.GetSysMenuBtnPage)
	e.POST("/sys_menu_btn/save", handle.SaveSysMenuBtn)
	e.POST("/sys_menu_btn/del", handle.DelSysMenuBtn)
}
