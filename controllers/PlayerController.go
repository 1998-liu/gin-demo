package controllers

import (
	"fmt"
	"gin-demo/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlayerController struct{}

type Player struct {
	Id          int    `json:"id"`
	Aid         int    `json:"aid" binding:"required"`
	Ref         string `json:"ref" binding:"required"`
	NickName    string `json:"nickName" binding:"required"`
	Declaration string `json:"declaration"`
	Avatar      string `json:"avatar"`
	Score       int    `json:"score"`
	AddTime     int    `json:"addTime"`
	UpdateTime  int    `json:"updateTime"`
}

// 获取参赛者
func (PlayerController) GetPlayers(c *gin.Context) {
	aidStr := c.DefaultQuery("aid", "0")
	aid, _ := strconv.Atoi(aidStr)
	result, err := models.Player{}.GetPlayers(aid)
	if err != nil {
		ReturnError(c, 4004, "没有相关信息")
		return
	}
	ReturnSuccess(c, 200, "success", result, int64(len(result)))
}

type playerResponse struct {
	Id       int    `json:"id"`
	Aid      int    `json:"aid"`
	Ref      string `json:"ref"`
	NickName string `json:"nickName"`
}

// 添加参赛者
func (PlayerController) AddPlayer(c *gin.Context) {
	var player Player
	err := c.BindJSON(&player)
	if err != nil {
		ReturnError(c, 4001, err.Error())
	}
	//查询活动是否存在
	activity, _ := models.Activity{}.GETActivityById(player.Aid)
	if activity.Id == 0 {
		ReturnError(c, 4001, "活动不存在")
		return
	}
	id, err := models.Player{}.AddPlayer(player.Aid, player.Ref, player.NickName, player.Declaration, player.Avatar)
	if id == 0 {
		ReturnError(c, 4001, fmt.Sprintf("添加失败，失败原因：%s", err.Error()))
		return
	}
	res := playerResponse{
		Id:       id,
		Aid:      player.Aid,
		Ref:      player.Ref,
		NickName: player.NickName,
	}
	ReturnSuccess(c, 200, "success", res, 1)
}
