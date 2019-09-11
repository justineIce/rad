package router

import (
	"{{.PackageName}}/manage-api/handle"
	"github.com/labstack/echo"
)

func SysRoleMenuRouter(e *echo.Group) {
	e.POST("/sys_role_menu/get", handle.GetSysRoleMenu)
	e.POST("/sys_role_menu/all", handle.GetSysRoleMenuAll)
	e.POST("/sys_role_menu/page", handle.GetSysRoleMenuPage)
	e.POST("/sys_role_menu/save", handle.SaveSysRoleMenu)
	e.POST("/sys_role_menu/del", handle.DelSysRoleMenu)
}
