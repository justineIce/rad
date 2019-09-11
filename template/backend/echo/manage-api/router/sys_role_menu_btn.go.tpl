package router

import (
	"{{.PackageName}}/manage-api/handle"
	"github.com/labstack/echo"
)

func SysRoleMenuBtnRouter(e *echo.Group) {
	e.POST("/sys_role_menu_btn/get", handle.GetSysRoleMenuBtn)
	e.POST("/sys_role_menu_btn/all", handle.GetSysRoleMenuBtnAll)
	e.POST("/sys_role_menu_btn/page", handle.GetSysRoleMenuBtnPage)
	e.POST("/sys_role_menu_btn/save", handle.SaveSysRoleMenuBtn)
	e.POST("/sys_role_menu_btn/del", handle.DelSysRoleMenuBtn)
}
