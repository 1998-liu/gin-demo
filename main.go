package main

import (
	"fmt"
	"gin-demo/routers"
	"gin-demo/utils"

	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath("config")
	viper.SetConfigFile("init")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

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
	host := viper.GetString("main.host")
	port := viper.GetString("main.port")
	fmt.Println("debug: ", host+":"+port)
	router.Run("127.0.0.1:8080")
}
