package main

import (
	"gin-demo/routers"
	"gin-demo/utils"
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
	router.Run("127.0.0.1:8080")
}
