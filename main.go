package main

import (
	"fmt"
	"gin-demo/config"
	_ "gin-demo/config"
	"gin-demo/routers"
	//"gin-demo/utils"
)

func main() {
	//初始化数据库
	//utils.InitMysql()
	router := routers.InitRouter()
	host := config.Conf.GetString("main.host")
	port := config.Conf.GetString("main.port")
	fmt.Println("debug: ", host+":"+port)
	router.Run(host + ":" + port)
}
