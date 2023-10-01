package models

import (
	"gin-demo/dao"
	"time"
)

type User struct {
	Id         int    `gorm:"column:id"`
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	AddTime    int64  `gorm:"column:addTime"`
	UpdateTime int64  `gorm:"column:updateTime"`
}

func (User) TableName() string {
	return "user"
}

func (User) GetUserByUsername(username string) (User, error) {
	var user User
	err := dao.Db.Where("username=?", username).First(&user).Error
	return user, err
}

func (User) GetUserById(uid int) (User, error) {
	var user User
	err := dao.Db.Where("id=?", uid).First(&user).Error
	return user, err
}

func (User) AddUser(username, password string) (int, error) {
	user := User{
		Username:   username,
		Password:   password,
		AddTime:    time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	err := dao.Db.Create(&user).Error
	return user.Id, err
}
