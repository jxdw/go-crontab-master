package main

import (
	"fmt"
	"os/exec"
)

var (
	cmd    *exec.Cmd
	output []byte
	err    error
)

func main() {
	//生成Cmd
	cmd = exec.Command("/bin/bash", "-c", "sleep 5; ls -l")

	//执行命令，捕获子进程的输出(pipe)
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}

	//打印子进程的输出
	fmt.Println(output)
	fmt.Println(string(output))
}
