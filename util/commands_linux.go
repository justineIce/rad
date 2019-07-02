package utils

import (
	"os/exec"
)

type Command struct {
	name string
	args []string
}

func Cmd(c Command) error {
	cmd := exec.Command(c.name, c.args...)
	err := cmd.Start()
	if err != nil {
		return err
	}
	cmd.Output()
	return nil
}


func Copy(source, target string) error {
	c := Command{name: "cp", args: []string{"-r",source, target}}
	return Cmd(c)
}
