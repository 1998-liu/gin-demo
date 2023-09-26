package controllers

import (
	"github.com/gin-gonic/gin"
)

type OrderController struct{}

type Search struct {
	Id      int    `json:"id"`
	OrderNo string `json:"order_no"`
}

func (o OrderController) GetList(c *gin.Context) {
	//接收 form 表单参数
	// id := c.PostForm("id")
	// orderNo := c.DefaultPostForm("order_no", "order123")
	// ReturnSuccess(c, 200, "success", "id:"+id+","+"orderNo:"+orderNo, 2)

	//使用 map 接收json参数
	// param := make(map[string]interface{}, 5)
	// err := c.BindJSON(&param)

	//使用结构体接收 json 参数
	param := &Search{}
	err := c.BindJSON(&param)
	if err == nil {
		ReturnSuccess(c, 200, "success", param, 1)
		return
	}
	ReturnError(c, 4001, gin.H{"err": err})
}
