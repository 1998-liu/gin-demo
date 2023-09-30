package models

import (
	"gin-demo/dao"
	"time"
)

type Activity struct {
	Id      int    `gorm:"column:id"`
	Name    string `gorm:"column:name"`
	AddTime int64  `gorm:"column:addTime"`
}

func (Activity) TableName() string {
	return "activity"
}

// 添加活动
func (Activity) AddActivity(name string) (int, error) {
	activity := Activity{
		Name:    name,
		AddTime: time.Now().Unix(),
	}
	err := dao.Db.Create(&activity).Error
	return activity.Id, err
}

// 通过活动名获取活动记录
func (Activity) GETActivityByName(name string) (Activity, error) {
	var activity Activity
	err := dao.Db.Where("name=?", name).First(&activity).Error
	return activity, err
}

// 通过活动 id 获取活动记录
func (Activity) GETActivityById(id int) (Activity, error) {
	var activity Activity
	err := dao.Db.Where("id=?", id).First(&activity).Error
	return activity, err
}

// 删除活动
func (Activity) DelActivityById(activity Activity) error {
	err := dao.Db.Delete(&activity).Error
	return err
}
