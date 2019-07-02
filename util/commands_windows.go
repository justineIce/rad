// +build !linux

package util

import (
	"os/exec"
	"strings"
)

const (
	WINDOWS = "windows"
	LINUX   = "linux"
)

type Command struct {
	name string
	args []string
}

func Copy(source, target string) (out []byte, err error) {
	args := []string{strings.ReplaceAll(source, "/", "\\"), strings.ReplaceAll(target, "/", "\\"), "/E", "/Y", "/D","/Q","/I"}
	out, err = exec.Command("xcopy", args...).Output()
	return
}
