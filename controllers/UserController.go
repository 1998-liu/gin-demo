package controllers

import "github.com/gin-gonic/gin"

type UserController struct{}

func (u UserController) GetUserInfo(c *gin.Context) {
	ReturnSuccess(c, 0, "success", "user_info", 1)
}

func (u UserController) GetUserList(c *gin.Context) {
	ReturnError(c, 4004, "获取列表错误")
}
