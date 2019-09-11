package router

import (
	"{{.PackageName}}/manage-api/handle"
	"github.com/labstack/echo"
)

func SysUserRouter(e *echo.Group) {
	e.POST("/sys_user/get", handle.GetSysUser)
	e.POST("/sys_user/all", handle.GetSysUserAll)
	e.POST("/sys_user/page", handle.GetSysUserPage)
	e.POST("/sys_user/save", handle.SaveSysUser)
	e.POST("/sys_user/del", handle.DelSysUser)
	e.POST("/sys_user/pwd/reset", handle.ResetPassword)
}
