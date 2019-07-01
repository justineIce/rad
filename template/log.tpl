package utils

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/zxbit2011/raddemo/config"
)

func NewLog(c *config.Log) *logs.BeeLogger {
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
