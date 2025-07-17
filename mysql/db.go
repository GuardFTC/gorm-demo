// Package mysql @Author:冯铁城 [17615007230@163.com] 2023-09-14 09:22:28
package mysql

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 数据库连接参数
var ip = "127.0.0.1"
var port = "3306"
var username = "root"
var password = "root"
var database = "golang_test"
var dsnConfig = "charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"

// InitDB 初始化数据库
func InitDB() *gorm.DB {

	//1.拼接DSN
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?%v", username, password, ip, port, database, dsnConfig)

	//2.打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("database connection error")
	}

	//3.返回数据库连接
	log.Printf("database connection success")
	return db
}
