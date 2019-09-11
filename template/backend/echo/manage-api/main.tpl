`package main

import (
	"context"
	"fmt"
	"{{.PackageName}}/manage-api/global"
	"{{.PackageName}}/manage-api/handle"
	"{{.PackageName}}/manage-api/router"
	"{{.PackageName}}/utils"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"os"
	"os/signal"
	"sort"
	"strings"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			println("异常：%v", err)
			time.Sleep(1 * time.Minute)
		}
	}()
	err := global.InitGlobal(&global.ConfPath{
		ConfigPath: "config/config.toml",
	})
	if err != nil {
		println("初始化失败：", err.Error())
		time.Sleep(1 * time.Minute)
		return
	}

	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: DefaultSkipper,
		Level:   5,
	}))
	e.Use(middleware.Secure())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("16a2t6WKgSlsdf75N40SFOhglShSJfg0ua8OnZ"))))
	e.Use(middleware.BodyDumpWithConfig(DefaultBodyDumpConfig))

    //上传文件访问
    e.Static("/files", "files")
    //上传的文件需登录方可访问
    e.GET("/files/*/auth/*", func(c echo.Context) error {
        return c.File(strings.TrimLeft(c.Request().URL.Path, "/"))
    }, handle.Filter)
	api := e.Group("/api")
	{
        api.GET("/img/code", handle.VerificationCode)
        //账号登录
        api.POST("/login", handle.Login)
        //账号登出
        api.POST("/logout", handle.LogOut)
        auth := api.Group("/auth", handle.Filter)
        {
            //账号菜单权限
            auth.POST("/login/menu", handle.GetLoginMenu)
            auth.POST("/files/upload/:auth", handle.UploadFile)
            auth.POST("/files/remove", handle.DelFileLog)
			router.FileLogRouter(auth)
			router.SysCompanyRouter(auth)
			router.SysDictRouter(auth)
			router.SysLogRouter(auth)
			router.SysMenuRouter(auth)
			router.SysMenuBtnRouter(auth)
			router.SysOfficeRouter(auth)
			router.SysRoleRouter(auth)
			router.SysRoleMenuRouter(auth)
			router.SysRoleMenuBtnRouter(auth)
			router.SysUserRouter(auth)
			router.SysUserRoleRouter(auth)
            // ===== router start =====
            {{range .Routers}}{{.}}
            {{end}}
            // ===== router end =====
        }
    }

	//路由打印
	var paths []string
	for _, v := range e.Routes() {
		if !strings.Contains(v.Name, "func1") && !strings.HasSuffix(v.Path, "/lists") &&
			!strings.HasSuffix(v.Path, "/history") &&
			!strings.HasSuffix(v.Path, "/list") &&
			!strings.HasSuffix(v.Path, "/classify") &&
			!strings.HasSuffix(v.Path, "/today") &&
			!strings.HasSuffix(v.Path, "/item") &&
			!strings.HasSuffix(v.Path, "/parentId") &&
			!strings.HasSuffix(v.Path, "/preview") &&
			!strings.HasSuffix(v.Path, "/resource") &&
			!strings.HasSuffix(v.Path, "/ids") &&
			!strings.HasSuffix(v.Path, "/node") &&
			!strings.HasSuffix(v.Path, "/detail") &&
			!strings.HasSuffix(v.Path, "/all") &&
			!strings.HasSuffix(v.Path, "/get") {
			paths = append(paths, v.Path)
		}
	}
	sort.Strings(paths)
	for _, value := range paths {
		fmt.Println("\"" + value + "\":\"\"")
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

//排除动态资源，如：图形验证码
func DefaultSkipper(c echo.Context) bool {
	if c.Path() == "/api/img/code" {
		return true
	}
	return false
}

var DefaultBodyDumpConfig = middleware.BodyDumpConfig{
	Skipper: BodyDumpDefaultSkipper,
	Handler: func(c echo.Context, reqBody, resBody []byte) {
		if len(resBody) <= 0 {
			return
		}
		result := utils.ToResultParam(resBody)
		if result.Ret <= 0 {
			return
		}
		//保存日志
		handle.SaveSysLog(c, result)
	},
}

//排除文件
func BodyDumpDefaultSkipper(c echo.Context) bool {
	/*if strings.Contains(c.Path(), "/api/m/files/") {
		return true
	}*/
	return false
}`
