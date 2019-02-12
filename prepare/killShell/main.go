package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err    error
	output []byte
}

var (
	ctx        context.Context
	cancelFunc context.CancelFunc
	cmd        *exec.Cmd
	resultChan chan *result
	res        *result
)

func main() {
	//执行1个cmd,让它在一个协程里去执行，让它执行2秒,1秒的时候，我们杀死cmd
	//sleep 2; echo hello;

	//context 有一个chan byte
	//cancelFunc:  关闭 close(chan byte)
	ctx, cancelFunc = context.WithCancel(context.TODO())

	//创建一个结果队列
	resultChan = make(chan *result, 1000)
	go func() {
		var (
			output []byte
			err    error
		)
		cmd = exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 5; echo hello;")
		//signal: killed

		//cmd = exec.CommandContext(ctx,"/bin/bash","-c"," echo hello;")
		//<nil> hello

		//select {case <- ctx.Done(): }
		//kill pid,进程ID,杀死子进程
		output, err = cmd.CombinedOutput()

		//把任务输出结果，传给main协程
		resultChan <- &result{
			err:    err,
			output: output,
		}
	}()

	//继续往下走
	time.Sleep(1 * time.Second)

	//取消上下文
	cancelFunc()

	//在main协程里，等待子协程的退出，并打印任务执行结果
	res = <-resultChan

	//打印任务执行结果
	fmt.Println(res.err, string(res.output))
}
