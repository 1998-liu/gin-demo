package controllers

import (
	"gin-demo/cache"
	"gin-demo/dao"
	"gin-demo/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VoteController struct{}

func (VoteController) AddVote(c *gin.Context) {
	uidStr := c.DefaultQuery("userId", "0")
	pidStr := c.DefaultQuery("playerId", "0")
	uid, _ := strconv.Atoi(uidStr)
	pid, _ := strconv.Atoi(pidStr)
	if uid == 0 || pid == 0 {
		ReturnError(c, 4001, "参数错误")
		return
	}
	user, _ := models.User{}.GetUserById(uid)
	if user.Id == 0 {
		ReturnError(c, 4001, "投票用户不存在")
		return
	}
	player, _ := models.Player{}.GetPlayerById(pid)
	if player.Id == 0 {
		ReturnError(c, 4001, "参赛者不存在")
		return
	}
	//开启事务
	tx := dao.Db.Begin()
	//每个投票用户投给参赛者只能投一次
	vote, _ := models.Vote{}.GetVoteInfo(uid, pid)
	if vote.Id != 0 {
		ReturnError(c, 4001, "已投票")
		return
	}
	res, err1 := models.Vote{}.AddVote(uid, pid, tx)
	//投票成功后，需要修改参赛者票数
	err2 := models.Player{}.UpdatePlayerScore(pid, tx)
	if err1 == nil && err2 == nil {
		ReturnSuccess(c, 200, "投票成功", res, 1)
		//提交事务
		tx.Commit()
		//更新 redis 排行榜缓存
		var redisKey string
		redisKey = "ranking:" + strconv.Itoa(player.Aid)
		cache.Rdb.ZIncrBy(cache.Rctx, redisKey, 1, pidStr)
		return
	}
	//回滚事务
	tx.Rollback()
	ReturnError(c, 4004, "投票失败，请联系管理员")
}
