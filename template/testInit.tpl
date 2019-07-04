package test

import "{{.PackageName}}/api/global"

func init() {
	global.InitGlobal(&global.ConfPath{
		ConfigPath: "../config/config.toml",
	})
}
