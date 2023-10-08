package controllers

import (
	"fmt"
	"gin-demo/cache"
	"gin-demo/models"
	"strconv"
	"time"

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
	result, err := models.Player{}.GetPlayers(aid, "id desc")
	if err != nil {
		ReturnError(c, 4004, "没有相关信息")
		return
	}
	ReturnSuccess(c, 200, "success", result, int64(len(result)))
}

// 获取排行榜
func (PlayerController) GetRanking(c *gin.Context) {
	aidStr := c.DefaultQuery("aid", "0")
	aid, _ := strconv.Atoi(aidStr)

	var redisKey string
	redisKey = "ranking:" + aidStr
	rs, err := cache.Rdb.ZRevRange(cache.Rctx, redisKey, 0, -1).Result()
	//判断缓存是否存在
	if err == nil && len(rs) > 0 {
		var players []models.Player
		for _, value := range rs {
			id, _ := strconv.Atoi(value)
			//缓存存在，通过参赛者 id 查询其信息
			rsInfo, _ := models.Player{}.GetPlayerById(id)
			if rsInfo.Id > 0 {
				players = append(players, rsInfo)
			}
		}
		count := int64(len(players))
		ReturnSuccess(c, 200, "success", players, count)
		return
	}
	rsDb, errDb := models.Player{}.GetPlayers(aid, "score desc")
	//缓存不存在，查询数据库并加入缓存
	if errDb == nil {
		for _, value := range rsDb {
			cache.Rdb.ZAdd(cache.Rctx, redisKey, cache.Zscore(value.Id, value.Score)).Err()
		}
		count := int64(len(rsDb))
		//设置过期时间
		cache.Rdb.Expire(cache.Rctx, redisKey, 24*time.Hour)
		ReturnSuccess(c, 200, "success", rsDb, count)
		return
	}
	ReturnError(c, 4004, "没有相关信息")
	return

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
