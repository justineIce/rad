package global

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"{{.PackageName}}/utils/config"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

var (
	//DB 数据库操作
	DB *gorm.DB
	//Conf 配置相关
	Conf *config.Config
	//Log 日志
	Log *logs.BeeLogger
	//RD redis
	RD *config.RedisConnPool
	//Session
	Session = func(c echo.Context) *config.USession { return config.GetSession(c) }
	//Secret
	Secret  = []byte("5eF6Xj8z#pZxBOkavlcPq^MmC09*S*!8!8Jor8V7*m0F3*zWReV%o3taoH%DI@ni")
	VerificationCode = "VerificationCode"
	VerificationNum = "VerificationNum"
)

//ConfPath 配置文件路径
type ConfPath struct {
	ConfigPath string
}

//InitGlobal 初始化各项配置
func InitGlobal(confPath *ConfPath) (err error) {
	// 初始化配置文件
	Conf, err = config.LoadGlobalConfig(confPath.ConfigPath)
	if err != nil {
		err = errors.New(fmt.Sprintf("初始化toml配置文件失败，error：%s", err.Error()))
		return
	}
	//初始化数据库连接
	DB, err = config.NewGorm(Conf)
	if err != nil {
		err = errors.New(fmt.Sprintf("初始化数据库失败，error：%s", err.Error()))
		return
	}
	//初始化redis
	RD, err = config.InitRedis(Conf.Redis, Conf.Project.Name)
	if err != nil {
		err = errors.New(fmt.Sprintf("初始化Redis数据库失败，error：%s", err.Error()))
		return
	}
	Log = config.InitLog(Conf.Log)
	return
}
