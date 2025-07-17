// Package mysql @Author:冯铁城 [17615007230@163.com] 2023-10-11 14:32:40
package mysql

import (
	"log"
	"time"

	"gorm-demo/mysql/model"
	"gorm.io/gorm"
)

// Update 更新数据
func Update(db *gorm.DB) error {

	//1.获取夏洛
	var student *mysql.Student
	db.Where("name = ?", "夏洛").First(&student)

	//2.单个全量更新
	//更新
	student.Age = 66
	student.UpdateTime = time.Now().Local()
	db.Save(&student)

	//查询验证
	updateStudent := &mysql.Student{ID: student.ID}
	db.Select("id", "age", "update_time").First(&updateStudent)
	log.Printf("更新成功 id=[%v] student.age=[%v] student.update_time=[%v]", updateStudent.ID, updateStudent.Age, updateStudent.UpdateTime)

	//3.单个更新单列
	//更新
	db.Model(&student).Update("email", "xialuo_update@xltfn.cn")

	//查询验证
	db.Select("id", "age", "email", "update_time").First(&updateStudent)
	log.Printf("更新成功 id=[%v] student.email=[%v] student.age=[%v] student.update_time=[%v]", updateStudent.ID, updateStudent.Email, updateStudent.Age, updateStudent.UpdateTime)

	//4.批量更新单列
	//更新
	db.Model(&mysql.Student{}).Where("age > ?", 18).Update("create_time", time.Now().Local())

	//查询验证
	var students []*mysql.Student
	db.Select("id", "create_time").Where("age > ?", 18).Find(&students)
	log.Println("批量更新成功")
	for _, student := range students {
		log.Printf("id = [%v] student.createTime=[%+v]", student.ID, student.CreateTime)
	}

	//5.批量更新多列
	//更新
	db.Model(&mysql.Student{}).Where("age <= ?", 18).Updates(&mysql.Student{UpdateTime: time.Now().Local(), Class: 11})

	//查询验证
	db.Select("id", "update_time", "class").Where("age <= ?", 18).Find(&students)
	log.Println("批量更新成功")
	for _, student := range students {
		log.Printf("id = [%v] student.update_time=[%+v] student.class=[%+v]", student.ID, student.UpdateTime, student.Class)
	}

	//6.不带条件的全部更新
	//更新
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&mysql.Teacher{}).Update("create_time", time.Now().Local())

	//查询验证
	var teachers []*mysql.Teacher
	db.Select("id", "create_time").Find(&teachers)
	log.Println("批量更新成功")
	for _, teacher := range teachers {
		log.Printf("id = [%v] teacher.create_time=[%+v]", teacher.ID, teacher.CreateTime)
	}

	//7.默认返回
	return nil
}
