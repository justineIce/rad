package router

import (
	"{{.PackageName}}/manage-api/handle"
	"github.com/labstack/echo"
)

func FileLogRouter(e *echo.Group) {
	e.POST("/file_log/get", handle.GetFileLog)
	e.POST("/file_log/all", handle.GetFileLogAll)
	e.POST("/file_log/page", handle.GetFileLogPage)
	e.POST("/file_log/save", handle.SaveFileLog)
	e.POST("/file_log/del", handle.DelFileLog)
}
