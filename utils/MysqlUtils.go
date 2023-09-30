package utils

import (
	"database/sql"
	"fmt"
	"gin-demo/config"
	"log"

	// 这个不加在编译的时候不会报错，但是在运行的时候就会报错,因为在编译的时候不需要用所以前面加_
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// 连接数据库
func InitMysql() {
	fmt.Println("InitMysql...")
	if db == nil {
		conf := config.Conf
		username := conf.GetString("mysql.username")
		password := conf.GetString("mysql.password")
		host := conf.GetString("mysql.host")
		port := conf.GetString("mysql.port")
		database := conf.GetString("mysql.database")
		suffix := conf.GetString("mysql.suffix")
		// db, _ = sql.Open("mysql", "root:jdtlh@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true&loc=Local")
		db, _ = sql.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+database+"?"+suffix)
		CreateTableWithUser()
	}
}

// 创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64) NOT NULL COMMENT "用户名",
		password VARCHAR(64) COMMENT "md5加密",
		status TINYINT(4) COMMENT "状态：0 正常，1 删除",
		createTime INT(10) COMMENT "创建时间"
	);`
	ModifyDB(sql)
}

// 原生sql操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

// 查询数据库
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}
