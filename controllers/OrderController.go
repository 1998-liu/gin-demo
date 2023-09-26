package controllers

import "github.com/gin-gonic/gin"

type OrderController struct{}

func (o OrderController) GetList(c *gin.Context) {
	ReturnSuccess(c, 200, "success", "order_list", 2)
}
