package main

import (
	"runtime"
	"crontab/master"
	"fmt"
	"flag"
	"time"
)

var (
	confFile string
)

func initArgs()  {
	flag.StringVar(&confFile,"config","./master.json","指定master.json")
	flag.Parse()
}

func initEnv()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
}


func main()  {

	var(
		err error
	)
	initArgs()
	//初始化线程
	initEnv()

	//启动API HTTP服务

	if err = master.InitConfig(confFile);err!=nil{
		goto ERR
	}


	//  任务管理器
	if err = master.InitJobMgr(); err != nil {
		goto ERR
	}

	if err = master.InitApiServer();err != nil{
		goto ERR
	}

	for{
		time.Sleep(1*time.Second)
	}

	return

ERR:
	fmt.Println(err)
}
