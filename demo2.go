package main

import (
	"fmt"
	"os/exec"
)

/**
* @Author: super
* @Date: 2021-02-01 10:59
* @Description: 获取执行输出
**/

func main() {
	var (
		cmd *exec.Cmd
		output []byte
		err error
	)

	// 生成Cmd
	cmd = exec.Command("/bin/bash", "-c", "ls -l")

	// 执行了命令, 捕获了子进程的输出( pipe )
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}

	// 打印子进程的输出
	fmt.Println(string(output))
}