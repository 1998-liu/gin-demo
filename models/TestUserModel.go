package models

import (
	"fmt"
	"gin-demo/utils"
)

type UserTest struct {
	Id         int
	Username   string
	Password   string
	Status     int //0:正常状态，1：删除
	CreateTime int64
}

// 插入
func InsertUser(user UserTest) (int64, error) {
	return utils.ModifyDB("insert into users(username,password,status,createTime) values (?,?,?,?)",
		user.Username, user.Password, user.Status, user.CreateTime)
}

// 按条件查询
func QueryUserWithCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println("sql")
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

// 根据用户名查询id
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWithCon(sql)
}

// 根据用户名和密码，查询id
func QueryUserWithParam(username, password string) int {
	sql := fmt.Sprintf("where username='%s' and password='%s'", username, password)
	return QueryUserWithCon(sql)
}
