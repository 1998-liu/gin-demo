package routers

import (
	"gin-demo/common/middleware/logger"
	"gin-demo/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	//使用日志中间件
	router.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	router.Use(logger.Recover)

	//注册测试
	router.POST("/register", controllers.RegisterPost)

	//路由组
	user := router.Group("/user")
	{
        user.POST("/register", controllers.UserController{}.Register)
		user.GET("/info", controllers.UserController{}.GetUserInfo)
		user.POST("/list", controllers.UserController{}.GetUserList)
		user.PUT("/add", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "user add")
		})
	}

	order := router.Group("/order")
	{
		order.POST("/list", controllers.OrderController{}.GetList)
		order.GET("/test", controllers.OrderController{}.Test)
	}
	return router
}
