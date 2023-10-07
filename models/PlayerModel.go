package models

import (
	"gin-demo/dao"
	"time"

	"github.com/jinzhu/gorm"
)

type Player struct {
	Id          int    `gorm:"column:id" json:"id"`
	Aid         int    `gorm:"column:aid" json:"aid"`
	Ref         string `gorm:"column:ref" json:"ref"`
	NickName    string `gorm:"column:nickName" json:"nickName"`
	Declaration string `gorm:"column:declaration" json:"declaration"`
	Avatar      string `gorm:"column:avatar" json:"avatar"`
	Score       int    `gorm:"column:score" json:"score"`
	AddTime     int    `gorm:"column:addTime" json:"addTime"`
	UpdateTime  int    `gorm:"column:updateTime" json:"updateTime"`
}

func (Player) GetTableName() string {
	return "player"
}

func (Player) GetPlayers(aid int, sort string) ([]Player, error) {
	var player []Player
	err := dao.Db.Where("aid=?", aid).Order(sort).Find(&player).Error
	return player, err
}

func (Player) GetPlayerById(pid int) (Player, error) {
	var player Player
	err := dao.Db.Where("id=?", pid).First(&player).Error
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

// 更新参赛者票数
func (Player) UpdatePlayerScore(id int, db *gorm.DB) error {
	var player Player
	err := db.Model(&player).Where("id=?", id).UpdateColumn("score", gorm.Expr("score + ?", 1)).Error
	return err
}
