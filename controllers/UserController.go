package controllers

import (
	"fmt"
	"gin-demo/models"
	"gin-demo/utils"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

type UserRegister struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Repassword string `json:"repassword" binding:"required,eqfield=Password"`
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 用户注册
func (u UserController) Register(c *gin.Context) {
	//接收用户名 密码 确认密码
	var up UserRegister
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
	_, err = models.User{}.AddUser(up.Username, utils.MD5(up.Password))
	if err != nil {
		ReturnError(c, 4001, fmt.Sprintf("注册失败，请联系管理员，错误原因：%s", err.Error()))
		return
	}
	ReturnSuccess(c, 200, "success", "注册成功", 1)
}

type UserApi struct {
    Id int `json:"id"`
    Username string `json:"username"`
}

// 用户登录
func (u UserController) Login(c *gin.Context) {
	var ul UserLogin
	if err := c.ShouldBindJSON(&ul); err != nil {
		ReturnError(c, 4001, err.Error())
        return
	}
	user, _ := models.User{}.GetUserByUsername(ul.Username)
    if user.Id == 0 {
        ReturnError(c, 4001, "用户名或者密码不正确")
        return
    }
    if user.Password != utils.MD5(ul.Password) {
        ReturnError(c, 4001, "用户名或者密码不正确")
        return
    }
    session := sessions.Default(c)
    session.Set("login:"+strconv.Itoa(user.Id), user.Id)
    session.Save()
    data := UserApi{Id: user.Id, Username: user.Username}
    ReturnSuccess(c, 200, "success", data, 1)
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
