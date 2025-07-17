// @Author:冯铁城 [17615007230@163.com] 2025-07-17 16:07:07
package main

import (
	"log"

	"gorm-demo/mysql"
	"gorm.io/gorm"
)

func main() {

	//1.链接数据库
	db := mysql.InitDB()

	//2.删除表
	if err := mysql.DropTable(db); err != nil {
		log.Fatal(err)
	}

	//3.创建表
	if err := mysql.CreateTable(db); err != nil {
		log.Fatal(err)
	}

	//4.开启事务-添加数据
	if err := db.Transaction(func(tx *gorm.DB) error {
		return mysql.Insert(tx)
	}); err != nil {
		log.Fatal(err)
	}

	//5.查询数据-初级查询
	if err := mysql.SimpleSelect(db); err != nil {
		log.Fatal(err)
	}

	//6.查询数据-条件查询
	if err := mysql.WhereSelect(db); err != nil {
		log.Fatal(err)
	}

	//7.查询数据-特殊查询
	if err := mysql.SpecialSelect(db); err != nil {
		log.Fatal(err)
	}

	//5.开启事务-更新数据
	if err := db.Transaction(func(tx *gorm.DB) error {
		return mysql.Update(tx)
	}); err != nil {
		log.Fatal(err)
	}
}
