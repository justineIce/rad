package router

import (
	"{{.PackageName}}/manage-api/handle"
	"github.com/labstack/echo"
)

func SysLogRouter(e *echo.Group) {
	e.POST("/sys_log/page", handle.GetSysLogPage)
}
