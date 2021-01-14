package util

import (
	"fmt"
	"os/exec"
	"strings"
)

//Command 命令结构体
type Command struct {
	CommandName string
	Params      []string
}

//ExecCommand 执行命令
func ExecCommand(command Command) error {
	cmd := exec.Command(command.CommandName, command.Params...)

	fmt.Println(strings.Join(cmd.Args, " "))

	if err := cmd.Start(); err != nil {
		return err
	}

	return nil
}
