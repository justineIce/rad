package router

import (
	"{{.PackageName}}/manage-api/handle"
	"github.com/labstack/echo"
)

func SysOfficeRouter(e *echo.Group) {
	e.POST("/sys_office/get", handle.GetSysOffice)
	e.POST("/sys_office/all", handle.GetSysOfficeAll)
	e.POST("/sys_office/page", handle.GetSysOfficePage)
	e.POST("/sys_office/save", handle.SaveSysOffice)
	e.POST("/sys_office/del", handle.DelSysOffice)
}
