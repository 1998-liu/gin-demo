package controllers

import (
	"fmt"
	"gin-demo/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

type UserParam struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Repassword string `json:"repassword" binding:"required,eqfield=Password"`
}

func (u UserController) Register(c *gin.Context) {
	//接收用户名 密码 确认密码
	var up UserParam
	if err := c.ShouldBindJSON(&up); err != nil {
		ReturnError(c, 4001, err.Error())
		return
	}
	//检查用户名是否已存在
	user, err := models.User{}.GetUserByUsername(up.Username)
	if user.Id != 0 {
		ReturnError(c, 4001, "用户名已存在")
		return
	}
	_, err = models.User{}.AddUser(up.Username, EncryMd5(up.Password))
	if err != nil {
		ReturnError(c, 4001, fmt.Sprintf("注册失败，请联系管理员，错误原因：%s", err.Error()))
		return
	}
	ReturnSuccess(c, 200, "success", "注册成功", 1)
}

// 测试接口
func (u UserController) GetUserInfo(c *gin.Context) {
	id := c.Query("id")
	name := c.DefaultQuery("name", "Gopher")
	userInfo := map[string]string{
		"id":   id,
		"name": name,
	}
	ReturnSuccess(c, 0, "success", userInfo, 1)
}

// 测试接口
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
