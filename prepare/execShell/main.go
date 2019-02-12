package main

import (
	"fmt"
	"os/exec"
)

var (
	cmd *exec.Cmd
	err error
)

func main() {
	//生成Cmd
	cmd = exec.Command("/bin/bash", "-c", "ls -al")

	//windows
	//cmd = exec.Command("C:\\cygwin64\\bin\\bash.exe","-c","echo1; echo2;")

	err = cmd.Run()
	if err != nil {
		fmt.Println("err:", err)
	}
}
