package main

import (
	"fmt"
	"os/exec"
)

/**
* @Author: super
* @Date: 2021-02-01 10:58
* @Description: 执行命令
**/

func main() {
	var (
		cmd *exec.Cmd
		err error
	)

	cmd = exec.Command("/bin/bash", "-c", "echo 1")

	err = cmd.Run()

	fmt.Println(err)
}
