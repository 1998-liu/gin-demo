package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) GetUserInfo(c *gin.Context) {
	id := c.Query("id")
	name := c.DefaultQuery("name", "Gopher")
	userInfo := map[string]string{
		"id":   id,
		"name": name,
	}
	ReturnSuccess(c, 0, "success", userInfo, 1)
}

func (u UserController) GetUserList(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	name := c.PostForm("name")
	userList := map[string]interface{}{
		"id":   id,
		"name": name,
	}
	// ReturnError(c, 4004, "获取列表错误")
	ReturnSuccess(c, 200, "success", userList, 1)
}
