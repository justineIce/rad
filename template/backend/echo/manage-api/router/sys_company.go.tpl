package router

import (
	"{{.PackageName}}/manage-api/handle"
	"github.com/labstack/echo"
)

func SysCompanyRouter(e *echo.Group) {
	e.POST("/sys_company/get", handle.GetSysCompany)
	e.POST("/sys_company/all", handle.GetSysCompanyAll)
	e.POST("/sys_company/page", handle.GetSysCompanyPage)
	e.POST("/sys_company/save", handle.SaveSysCompany)
	e.POST("/sys_company/del", handle.DelSysCompany)
}
