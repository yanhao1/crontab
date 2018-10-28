package main

import (
	"flag"
	"runtime"
	"crontab/worker"
	"fmt"
	"time"
)

var (
	confFile string
)
//解析命令参数
func initArgs()  {
	//worker --config  ./worker.json
	//worker -h 

	flag.StringVar(&confFile,"config","./worker.json","worker.json")
	flag.Parse()

}

//初始化线程数

func initEnv()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main(){
	var(
		err error
	)

	initArgs()

	initEnv()


	//加载配置

	if err = worker.InitConfig(confFile);err !=nil{
		goto ERR
	}

	if err = worker.InitJobMgr();err!=nil{
		goto ERR
	}

	for{
		time.Sleep(1*time.Second)
	}

	return

	ERR:
		fmt.Println(err)

}