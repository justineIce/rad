package test

import "{{.PackageName}}/global"

func init() {
	global.InitGlobal(&global.ConfPath{
		ConfigPath: "../config/config.toml",
	})
}
