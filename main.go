// @Author:冯铁城 [17615007230@163.com] 2025-07-17 16:07:07
package main

import (
	"gorm-demo/_gen"
	"gorm-demo/_gen/query"
	"gorm-demo/common"
	"gorm-demo/mysql"
	"gorm.io/gorm"
)

func main() {

	//1.链接数据库
	db := mysql.InitDB()

	//2.gorm测试
	//gormTest(db)

	//3.gen生成器测试
	//genGeneratorTest(db)

	//4.gen测试
	genTest(db)
}

// gorm测试
func gormTest(db *gorm.DB) {

	//1.删除表
	common.CheckError(mysql.DropTable(db))

	//2.创建表
	common.CheckError(mysql.CreateTable(db))

	//3.开启事务-添加数据
	common.CheckGormTransactionError(db, mysql.Insert)

	//4.查询数据-初级查询
	common.CheckError(mysql.SimpleSelect(db))

	//5.查询数据-条件查询
	common.CheckError(mysql.WhereSelect(db))

	//6.查询数据-特殊查询
	common.CheckError(mysql.SpecialSelect(db))

	//7.开启事务-更新数据
	common.CheckGormTransactionError(db, mysql.Update)

	//8.开启事务-删除数据
	common.CheckGormTransactionError(db, mysql.Delete)
}

// gen生成器测试
func genGeneratorTest(db *gorm.DB) {
	_gen.Generate(db)
}

// gen测试
func genTest(db *gorm.DB) {

	//1.删除表
	common.CheckError(mysql.DropTable(db))

	//2.创建表
	common.CheckError(mysql.CreateTable(db))

	//3.初始化全局入口
	q := query.Use(db)

	//4.保存数据
	common.CheckGenTransactionError(q, _gen.Insert)
}
