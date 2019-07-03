package util

import (
	"os/exec"
	"strings"
)

func Copy(source, target, exclude string) (out []byte, err error) {
	var args []string
	args = []string{strings.ReplaceAll(source, "/", "\\"), strings.ReplaceAll(target, "/", "\\"),
		"/E", "/Y", "/D", "/Q", "/I"}
	if exclude != "" {
		args = append(args, "/EXCLUDE:"+strings.ReplaceAll(exclude, "/", "\\"))
	}
	out, err = exec.Command("xcopy", args...).Output()
	return
}
