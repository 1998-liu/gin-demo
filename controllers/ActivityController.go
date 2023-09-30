package controllers

import (
	"fmt"
	"gin-demo/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ActivityController struct{}

type Activity struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// 添加活动
func (ActivityController) AddActivity(c *gin.Context) {
	var activity Activity
	err := c.BindJSON(&activity)
	if activity.Name == "" {
		ReturnError(c, 4001, "活动名不能为空")
		return
	}
	//查看活动名是否重复
	a, _ := models.Activity{}.GETActivityByName(activity.Name)
	if a.Id != 0 {
		ReturnError(c, 4001, "活动名已重复")
		return
	}
	activity.Id, err = models.Activity{}.AddActivity(activity.Name)
	if err != nil {
		ReturnError(c, 4001, fmt.Sprintf("活动添加失败，失败原因：%s", err.Error()))
		return
	}
	ReturnSuccess(c, 200, "success", activity, 1)
}

// 删除活动
func (ActivityController) DelActivity(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		ReturnError(c, 4001, "参数错误")
		return
	}
	aid, _ := strconv.Atoi(id)
	//检查活动是否存在
	activity, _ := models.Activity{}.GETActivityById(aid)
	if activity.Id == 0 {
		ReturnError(c, 4001, "活动不存在，删除失败")
		return
	}
	err := models.Activity{}.DelActivityById(activity)
	if err != nil {
		ReturnError(c, 4001, fmt.Sprintf("活动删除失败，失败原因：%s", err.Error()))
		return
	}
	ReturnSuccess(c, 200, "success", "删除成功", 1)
}
