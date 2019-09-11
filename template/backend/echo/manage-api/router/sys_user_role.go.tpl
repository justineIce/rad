package router

import (
	"{{.PackageName}}/manage-api/handle"
	"github.com/labstack/echo"
)

func SysUserRoleRouter(e *echo.Group) {
	e.POST("/sys_user_role/get", handle.GetSysUserRole)
	e.POST("/sys_user_role/all", handle.GetSysUserRoleAll)
	e.POST("/sys_user_role/page", handle.GetSysUserRolePage)
	e.POST("/sys_user_role/save", handle.SaveSysUserRole)
	e.POST("/sys_user_role/del", handle.DelSysUserRole)
}
