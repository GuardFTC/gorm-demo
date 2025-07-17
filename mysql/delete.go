// Package mysql @Author:冯铁城 [17615007230@163.com] 2023-10-11 14:57:42
package mysql

import (
	mysql "gorm-demo/mysql/model"
	"gorm.io/gorm"
)

func Delete(db *gorm.DB) {

	//1.查询老师数据
	var teacher *mysql.Teacher
	db.Where("name = ?", "王老师").First(&teacher)

	//2.根据主键ID删除老师数据
	db.Delete(&teacher)

	//3.根据名称删除批量删除
	db.Where("name = ?", "校长").Delete(&mysql.Teacher{})

	//4.批量查询学生数据
	var students []*mysql.Student
	db.Order("id desc").Find(&students)

	//5.根据主键ID删除学生数据
	db.Delete(&mysql.Student{}, students[0].ID)

	//6.根据主键ID批量删除学生数据
	db.Delete(&students)

	//7.中间表批量删除
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&mysql.TeacherAndStudent{})
}
