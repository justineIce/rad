`package main

import (
	"context"
	"fmt"
	"{{.PackageName}}/manage-api/global"
	"{{.PackageName}}/manage-api/handle"
	"{{.PackageName}}/manage-api/router"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"os"
	"os/signal"
	"strings"
	"time"
)

func main() {
	global.InitGlobal(&global.ConfPath{
		ConfigPath: "config/config.toml",
	})
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("16a2t6WKgSlsdf75N40SFOhglShSJfg0ua8OnZ"))))

	api := e.Group("/api")
	{
        api.GET("/img/code", handle.VerificationCode)
        //账号登录
        api.POST("/login", handle.Login)
        //上传文件访问
        e.Static("/files", "files")
        //上传的文件需登录方可访问
        e.GET("/files/*/auth/*", func(c echo.Context) error {
            return c.File(strings.TrimLeft(c.Request().URL.Path, "/"))
        }, handle.Filter)
        auth := api.Group("/auth", handle.Filter)
        {

            //账号登出
            auth.GET("/logout", handle.LogOut)
            //账号菜单权限
            auth.GET("/login/menu", handle.GetLoginMenu)
            auth.POST("/files/upload/:auth", handle.UploadFile)
            auth.POST("/files/remove", handle.DelFileLog)
            // ===== router start =====
            {{range .Routers}}{{.}}
            {{end}}
            // ===== router end =====
        }
    }
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", global.Conf.HTTP.Port)); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
`
