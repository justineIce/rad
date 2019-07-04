package config

import (
	"fmt"
	"github.com/astaxie/beego/logs"
)

// GormConfig Gorm配置文件
type LogConfig struct {
	Level    int
	Path     string
	MaxDays  int
	Separate string
}

func NewLog(c *LogConfig) *logs.BeeLogger {
	conf := fmt.Sprintf(
		`{
			"filename": "%s",
			"maxdays": %d,
			"daily": true,
			"rotate": true,
			"level": %d,
			"separate": "[%s]"
		}`,
		c.Path,
		c.MaxDays,
		c.Level,
		c.Separate,
	)
	logs.SetLogger(logs.AdapterMultiFile, conf)
	logs.SetLogger("console")
	logs.EnableFuncCallDepth(true)
	return logs.GetBeeLogger()
}
