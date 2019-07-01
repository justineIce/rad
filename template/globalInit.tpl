package global

import (
	"github.com/jinzhu/gorm"
	"{{.PackageName}}/config"
	"time"
	"gopkg.in/go-playground/validator.v9"
)
var (
	//DB 数据库操作
	DB *gorm.DB
	//Conf 配置相关
	Conf *config.Config
	//Validate struct验证
	Validate *validator.Validate
)

//ConfPath 配置文件路径
type ConfPath struct {
	ConfigPath string
}

//InitGlobal 初始化各项配置
func InitGlobal(confPath *ConfPath) {
	// 初始化配置文件
	err := config.LoadGlobalConfig(confPath.ConfigPath)
	if err != nil {
		panic(err)
	}

	//获取配置参数
	Conf = config.GetGlobalConfig()

	//初始化数据
	DB, err = NewGorm()
	if err != nil {
		panic(err)
	}
}

// NewGorm 创建DB实例
func NewGorm() (*gorm.DB, error) {
	db, err := gorm.Open(Conf.Gorm.DBType, Conf.MySQL.DSN())
	if err != nil {
		return nil, err
	}
	if Conf.Gorm.Debug {
		db = db.Debug()
	}
	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(Conf.Gorm.MaxIdleConns)
	db.DB().SetMaxOpenConns(Conf.Gorm.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(Conf.Gorm.MaxLifetime) * time.Second)
	return db, nil
}
