package main

import (
	"fmt"
	"gin-demo/config"
	_ "gin-demo/config"
	"gin-demo/routers"
	"gin-demo/utils"
	// "github.com/spf13/viper"
)

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	//     c.JSON(200, gin.H{
	//         "message": "pong",
	//     })
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080

	//初始化数据库
	utils.InitMysql()
	router := routers.InitRouter()
	host := config.Conf.GetString("main.host")
	port := config.Conf.GetString("main.port")
	fmt.Println("debug: ", host+":"+port)
	router.Run(host + ":" + port)
}
