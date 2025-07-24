// Package _gen @Author:冯铁城 [17615007230@163.com] 2025-07-23 14:53:52
package _gen

import (
	"gorm.io/gen"
	"gorm.io/gorm"
)

// Generate 生成代码
func Generate(db *gorm.DB) {

	//1.创建自动生成器
	generator := gen.NewGenerator(gen.Config{
		OutPath:      "_gen/query",
		ModelPkgPath: "model",
		Mode:         gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
	})

	//2.使用Gorm数据源
	generator.UseDB(db)

	//3.为数据库中所有表自动生成相关代码
	generator.ApplyBasic(generator.GenerateAllTable()...)

	//4.执行
	generator.Execute()
}
