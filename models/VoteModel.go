package models

import (
	"gin-demo/dao"
	"time"

	"github.com/jinzhu/gorm"
)

type Vote struct {
	Id       int `gorm:"column:id"`
	UserId   int `gorm:"column:userId"`
	PlayerId int `gorm:"column:playerId"`
	AddTime  int `gorm:"column:addTime"`
}

func (Vote) GetTableName() string {
	return "vote"
}

func (Vote) GetVoteInfo(uid, pid int) (Vote, error) {
	var vote Vote
	err := dao.Db.Where("userId=? and playerId=?", uid, pid).First(&vote).Error
	return vote, err
}

func (Vote) AddVote(uid, pid int, db *gorm.DB) (Vote, error) {
	vote := Vote{UserId: uid, PlayerId: pid, AddTime: int(time.Now().Unix())}
	err := db.Create(&vote).Error
	return vote, err
}
