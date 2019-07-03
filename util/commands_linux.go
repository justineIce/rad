// +build linux

package utils

import (
	"os/exec"
)

type Command struct {
	name string
	args []string
}

func Cmd(c Command) (out []byte, err error) {
	cmd := exec.Command(c.name, c.args...)
	err := cmd.Start()
	if err != nil {
		return err
	}
	return cmd.Output()
}

func Copy(source, target, exclude, exclude string) (out []byte, err error) {
	c := Command{name: "cp", args: []string{"-r", source, target}}
	return Cmd(c)
}
