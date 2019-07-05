package test

import "{{.PackageName}}/manage-api/global"

func init() {
	global.InitGlobal(&global.ConfPath{
		ConfigPath: "../config/config.toml",
	})
}
