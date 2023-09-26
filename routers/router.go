package routers

import (
	"gin-demo/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//注册
	router.POST("/register", controllers.RegisterPost)

	//路由组
	user := router.Group("/user")
	{
		user.GET("/info", controllers.UserController{}.GetUserInfo)
		user.POST("/list", controllers.UserController{}.GetUserList)
		user.PUT("/add", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "user add")
		})
	}

	order := router.Group("/order")
	{
		order.POST("/list", controllers.OrderController{}.GetList)
	}
	return router
}
