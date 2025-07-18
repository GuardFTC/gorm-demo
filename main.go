// @Author:冯铁城 [17615007230@163.com] 2025-07-17 16:07:07
package main

import (
	"gorm-demo/common"
	"gorm-demo/mysql"
)

func main() {

	//1.链接数据库
	db := mysql.InitDB()

	//2.删除表
	common.CheckError(mysql.DropTable(db))

	//3.创建表
	common.CheckError(mysql.CreateTable(db))

	//4.开启事务-添加数据
	common.CheckTransactionError(db, mysql.Insert)

	//5.查询数据-初级查询
	common.CheckError(mysql.SimpleSelect(db))

	//6.查询数据-条件查询
	common.CheckError(mysql.WhereSelect(db))

	//7.查询数据-特殊查询
	common.CheckError(mysql.SpecialSelect(db))

	//8.开启事务-更新数据
	common.CheckTransactionError(db, mysql.Update)

	//9.开启事务-删除数据
	common.CheckTransactionError(db, mysql.Delete)
}
