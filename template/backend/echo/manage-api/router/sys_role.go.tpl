package router

import (
	"{{.PackageName}}/manage-api/handle"
	"github.com/labstack/echo"
)

func SysRoleRouter(e *echo.Group) {
	e.POST("/sys_role/get", handle.GetSysRole)
	e.POST("/sys_role/all", handle.GetSysRoleAll)
	e.POST("/sys_role/page", handle.GetSysRolePage)
	e.POST("/sys_role/save", handle.SaveSysRole)
	e.POST("/sys_role/del", handle.DelSysRole)
}
