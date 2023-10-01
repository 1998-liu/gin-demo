package models

import (
	"gin-demo/dao"
	"time"
)

type Player struct {
	Id          int    `gorm:"column:id"`
	Aid         int    `gorm:"column:aid"`
	Ref         string `gorm:"column:ref"`
	NickName    string `gorm:"column:nickName"`
	Declaration string `gorm:"column:declaration"`
	Avatar      string `gorm:"column:avatar"`
	Score       int    `gorm:"column:score"`
	AddTime     int    `gorm:"column:addTime"`
	UpdateTime  int    `gorm:"column:updateTime"`
}

func (Player) GetTableName() string {
	return "player"
}

func (Player) GetPlayers(aid int) ([]Player, error) {
	var player []Player
	err := dao.Db.Where("aid=?", aid).Find(&player).Error
	return player, err
}

// 添加参赛者
func (Player) AddPlayer(aid int, ref, nickName, declaration, avatar string) (int, error) {
	var player = Player{
		Aid:         aid,
		Ref:         ref,
		NickName:    nickName,
		Declaration: declaration,
		Avatar:      avatar,
		AddTime:     int(time.Now().Unix()),
		UpdateTime:  int(time.Now().Unix()),
	}
	dao.Db.LogMode(true)
	err := dao.Db.Create(&player).Error
	return player.Id, err
}
