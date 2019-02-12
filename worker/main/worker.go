package main

import (
	"flag"
	"go-crontab-master/worker"
	"runtime"
	"go-crontab-master/common"
	"time"
)

var (
	confFile string //配置文件路径
)

//解析命令行参数
func initArgs()  {
	//worker -config ./worker.json
	//worker -h
	flag.StringVar(&confFile,"config","./worker.json","worker.json")
	flag.Parse()
}

//初始化线程数量
func initEnv()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main()  {
	var (
		err error
	)

	//初始化命令行参数
	initArgs()

	//初始化线程
	initEnv()

	//加载配置
	if err = worker.InitConfig(confFile); err != nil{
		common.FmtErr(err)
	}

	//服务注册
	if err = worker.InitRegister(); err != nil{
		common.FmtErr(err)
	}

	//启动日志协程
	if err = worker.InitLogSink(); err != nil{
		common.FmtErr(err)
	}

	//启动执行器
	if err = worker.InitLogSink(); err != nil{
		common.FmtErr(err)
	}

	//初始化任务管理器
	if err = worker.InitJobMgr(); err != nil{
		common.FmtErr(err)
	}

	//正常退出
	for {
		time.Sleep(1*time.Second)
	}
	return
}