// Package _gen @Author:冯铁城 [17615007230@163.com] 2025-07-24 11:16:42
package _gen

import (
	"gorm-demo/_gen/model"
	"gorm-demo/_gen/query"
	"gorm.io/gorm"
	"log"
)

// Update 更新数据
func Update(tx *query.Query) error {

	//1.单个数据 更新单列
	info, err := tx.Student.Where(tx.Student.ID.Eq(1)).UpdateSimple(tx.Student.Name.Value("夏洛_update"))
	if err != nil {
		return err
	} else {
		log.Printf("更新成功 受影响行数=[%v]", info.RowsAffected)
	}

	//2.单个数据 更新多列 结构体传参
	dbStudent := &model.Student{Age: 19, Email: "xiaolou_update@xltfn.cn"}
	updates, err := tx.Student.Where(tx.Student.ID.Eq(1)).Updates(dbStudent)
	if err != nil {
		return err
	} else {
		log.Printf("更新成功 受影响行数=[%v]", updates.RowsAffected)
	}

	//3.单个数据 更新多列 直接设置值
	simple, err := tx.Student.Where(tx.Student.ID.Eq(1)).UpdateSimple(tx.Student.Age.Add(1), tx.Student.Name.Value("夏洛_update_2"))
	if err != nil {
		return err
	} else {
		log.Printf("更新成功 受影响行数=[%v]", simple.RowsAffected)
	}

	//4.多个数据更新单例
	updateSimple, err := tx.Student.Where(tx.Student.Age.Gt(18)).UpdateSimple(tx.Student.Age.Add(10))
	if err != nil {
		return err
	} else {
		log.Printf("更新成功 受影响行数=[%v]", updateSimple.RowsAffected)
	}

	//5.多个数据更新多列
	updateData := map[string]any{
		"age":  gorm.Expr("age+1"),
		"name": gorm.Expr("concat(name,'_update')"),
	}
	resultInfo, err := tx.Student.Where(tx.Student.Age.Lte(18)).Updates(updateData)
	if err != nil {
		return err
	} else {
		log.Printf("更新成功 受影响行数=[%v]", resultInfo.RowsAffected)
	}

	//6.默认返回
	return nil
}
