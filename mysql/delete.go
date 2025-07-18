// Package mysql @Author:冯铁城 [17615007230@163.com] 2023-10-11 14:57:42
package mysql

import (
	"log"

	mysql "gorm-demo/mysql/model"
	"gorm.io/gorm"
)

func Delete(db *gorm.DB) error {

	//1.查询老师数据
	var teacher *mysql.Teacher
	db.Where("name = ?", "王老师").First(&teacher)

	//2.根据主键ID删除老师数据
	//删除
	if tx := db.Delete(&teacher); tx.Error != nil {
		return tx.Error
	}

	//查询验证
	var count int64
	db.Where("id-?", teacher.ID).Count(&count)
	if count == 0 {
		log.Printf("王老师已删除")
	}

	//3.根据条件-名称删除批量删除
	//删除
	if tx := db.Where("name = ?", "校长").Delete(&mysql.Teacher{}); tx.Error != nil {
		return tx.Error
	}

	//查询验证
	db.Where("name = ?", "校长").Count(&count)
	if count == 0 {
		log.Printf("校长已删除")
	}

	//4.批量查询学生数据
	var students []*mysql.Student
	db.Order("id desc").Find(&students)

	//5.根据主键ID删除学生数据
	//删除
	if tx := db.Delete(&mysql.Student{}, students[0].ID); tx.Error != nil {
		return tx.Error
	}

	//查询验证
	db.Where("id = ?", students[0].ID).Count(&count)
	if count == 0 {
		log.Printf("学生id=%v已删除", students[0].ID)
	}

	//6.根据主键ID批量删除学生数据
	//删除
	if tx := db.Delete(&students); tx.Error != nil {
		return tx.Error
	}

	//查询验证
	db.Model(&mysql.Student{}).Count(&count)
	if count == 0 {
		log.Printf("所有学生已删除")
	}

	//7.中间表批量删除
	//删除
	if tx := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&mysql.TeacherAndStudent{}); tx.Error != nil {
		return tx.Error
	}

	//查询验证
	db.Model(&mysql.TeacherAndStudent{}).Count(&count)
	if count == 0 {
		log.Printf("所有老师学生关系已删除")
	}

	//8.默认返回
	return nil
}
