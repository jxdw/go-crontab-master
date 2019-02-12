package main

import (
	"flag"
	"go-crontab-master/common"
	"go-crontab-master/master"
	"runtime"
	"time"
)

var (
	confFile string //配置文件路径
	err error
)

//解析命令行参数
func initArgs()  {
	// master -config ./master.json -xxx 123 -yyy ddd
	flag.StringVar(&confFile,"config","./master.json","指定master.json")
	flag.Parse()
}
//初始化线程数量
func initEnv()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func main()  {
	//初始化命令行参数
	initArgs()

	//初始化线程
	initEnv()

	//加载配置
	if err = master.InitConfig(confFile); err != nil{
		common.FmtErr(err)
	}

	//初始化服务发现模块
	if err = master.InitWorkerMgr(); err != nil{
		common.FmtErr(err)
	}

	//日志管理器
	if err = master.InitLogMgr(); err != nil{
		common.FmtErr(err)
	}

	//任务管理器
	if err = master.InitJobMgr(); err != nil{
		common.FmtErr(err)
	}

	//启动Api Http服务
	if err = master.InitApiServer(); err != nil{
		common.FmtErr(err)
	}

	//正常退出
	for{
		time.Sleep(1*time.Second)
	}
	return
}

