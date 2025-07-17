// Package mysql @Author:冯铁城 [17615007230@163.com] 2023-10-11 09:46:46
package mysql

import (
	"log"

	"gorm-demo/mysql/model"
	"gorm.io/gorm"
)

// SimpleSelect 初级查询
func SimpleSelect(db *gorm.DB) error {

	//1.查询全部学生信息,不携带任何条件
	var students []mysql.Student
	if result := db.Find(&students); result.Error != nil {
		return result.Error
	}
	log.Println("get all data:")
	for _, student := range students {
		log.Printf("student=[%+v]", student)
	}

	//2.获取主键正序排列的第一个学生
	var firstStudent mysql.Student
	if result := db.First(&firstStudent); result.Error != nil {
		return result.Error
	}
	log.Println("get order by primary key asc first data:")
	log.Printf("firstStudent=[%+v]", firstStudent)

	//3.获取ID倒序排列的第一个学生
	var lastStudent mysql.Student
	if result := db.Last(&lastStudent); result.Error != nil {
		return result.Error
	}
	log.Println("get order by primary key desc first data:")
	log.Printf("lastStudent=[%+v]", lastStudent)

	//4.根据主键ID查询学生
	var studentById mysql.Student
	if result := db.First(&studentById, lastStudent.ID); result.Error != nil {
		return result.Error
	}
	log.Println("get data by id:")
	log.Printf("studentById=[%+v]", studentById)

	//5.根据主键ID查询学生
	studentByIdAndHasId := mysql.Student{ID: lastStudent.ID}
	if result := db.First(&studentByIdAndHasId); result.Error != nil {
		return result.Error
	}
	log.Println("get data by id:")
	log.Printf("studentByIdAndHasId=[%+v]", studentByIdAndHasId)

	//6.根据主键ID批量查询学生
	var studentsByIds []mysql.Student
	if result := db.Find(&studentsByIds, []uint{firstStudent.ID, lastStudent.ID}); result.Error != nil {
		return result.Error
	}
	log.Println("get batch data by id:")
	for _, studentsById := range studentsByIds {
		log.Printf("studentsById=[%+v]", studentsById)
	}

	//7.默认返回
	return nil
}

// WhereSelect 条件查询
func WhereSelect(db *gorm.DB) error {

	//1.等值查询，获取符合条件的第一个对象
	var studentByNameFirst mysql.Student
	if result := db.Where("name = ?", "马冬梅").First(&studentByNameFirst); result.Error != nil {
		return result.Error
	}
	log.Println("get where equal first data:")
	log.Printf("studentByNameFirst=[%+v]", studentByNameFirst)

	//2.等值查询，获取符合条件的全部对象
	var studentsByGrade []mysql.Student
	if result := db.Where("grade = ?", 3).Find(&studentsByGrade); result.Error != nil {
		return result.Error
	}
	log.Println("get where equal all data:")
	for _, studentByGrade := range studentsByGrade {
		log.Printf("studentByGrade=[%+v]", studentByGrade)
	}

	//3.不等值查询，获取符合条件的全部对象
	var studentsByGradeNotEqual []*mysql.Student
	if result := db.Where("grade <> ?", 3).Find(&studentsByGradeNotEqual); result.Error != nil {
		return result.Error
	}
	log.Println("get where not equal all data:")
	for _, studentByGradeNotEqual := range studentsByGradeNotEqual {
		log.Printf("studentByGradeNotEqual=[%+v]", studentByGradeNotEqual)
	}

	//4.in查询，获取符合条件的全部对象
	var studentsIn []mysql.Student
	if result := db.Where("age in ?", []uint{17, 19}).Find(&studentsIn); result.Error != nil {
		return result.Error
	}
	log.Println("get where in all data:")
	for _, studentIn := range studentsIn {
		log.Printf("studentIn=[%+v]", studentIn)
	}

	//5.like查询，获取符合条件的全部对象
	var studentsLike []mysql.Student
	if result := db.Where("email like ?", "%zhang%").Find(&studentsLike); result.Error != nil {
		return result.Error
	}
	log.Println("get where like all data:")
	for _, studentLike := range studentsLike {
		log.Printf("studentLike=[%+v]", studentLike)
	}

	//6.大于/大于等于查询
	var studentsGqOrGe []mysql.Student
	if result := db.Where("age > ?", 18).Find(&studentsGqOrGe); result.Error != nil {
		return result.Error
	}
	log.Println("get where gq/ge all data:")
	for _, studentGqOrGe := range studentsGqOrGe {
		log.Printf("studentGqOrGe=[%+v]", studentGqOrGe)
	}

	//7.小于/小于等于查询
	var studentsLqOrLe []mysql.Student
	if result := db.Where("age <= ?", 18).Find(&studentsLqOrLe); result.Error != nil {
		return result.Error
	}
	log.Println("get where lq/le all data:")
	for _, studentLqOrLe := range studentsLqOrLe {
		log.Printf("studentLqOrLe=[%+v]", studentLqOrLe)
	}

	//8.AND查询
	var studentsAnd []mysql.Student
	if result := db.Where("age >= ? and name = ?", 18, "陈凯").Find(&studentsAnd); result.Error != nil {
		return result.Error
	}
	log.Println("get where and all data:")
	for _, studentAnd := range studentsAnd {
		log.Printf("studentAnd=[%+v]", studentAnd)
	}

	//9.OR查询
	var studentsOr []mysql.Student
	if result := db.Where("age > ?", 20).Or("name = ?", "马冬梅").Find(&studentsOr); result.Error != nil {
		return result.Error
	}
	log.Println("get where or all data:")
	for _, studentOr := range studentsOr {
		log.Printf("studentOr=[%+v]", studentOr)
	}

	//10.结构映射查询
	var students []mysql.Student
	if result := db.Where(mysql.Student{Age: 18, Name: "夏洛"}).Find(&students); result.Error != nil {
		return result.Error
	}
	log.Println("get where all data:")
	for _, student := range students {
		log.Printf("student=[%+v]", student)
	}

	//11.内联查询(不使用where，直接通过Find写查询条件)
	var inlineStudents []mysql.Student
	if result := db.Find(&inlineStudents, "age > ? and age <= ?", 17, 21); result.Error != nil {
		return result.Error
	}
	log.Println("get where inline all data:")
	for _, inlineStudent := range inlineStudents {
		log.Printf("inlineStudent=[%+v]", inlineStudent)
	}

	//12.NOT查询
	var notStudents []mysql.Student
	if result := db.Not("age in ?", []uint{17, 19}).Find(&notStudents); result.Error != nil {
		return result.Error
	}
	log.Println("get where not all data:")
	for _, notStudent := range notStudents {
		log.Printf("notStudent=[%+v]", notStudent)
	}

	//13.查询指定字段
	var selectStudents []mysql.Student
	if result := db.Select("id", "name").Where("age = ?", 18).Find(&selectStudents); result.Error != nil {
		return result.Error
	}
	log.Println("get where select all data:")
	for _, selectStudent := range selectStudents {
		log.Printf("selectStudent=[%+v]", selectStudent)
	}

	//14.查询数量
	var count int64
	if result := db.Model(&mysql.Student{}).Count(&count); result.Error != nil {
		return result.Error
	}
	log.Println("get count:")
	log.Printf("count=[%v]", count)

	//15.默认返回
	return nil
}

// group查询结果
type groupResult struct {
	Age   int
	Total int
}

// SpecialSelect 特殊查询
func SpecialSelect(db *gorm.DB) error {

	//1.排序，获取全部数据，根据年龄倒序，id正序
	var orderStudents []mysql.Student
	if result := db.Order("age desc,id asc").Find(&orderStudents); result.Error != nil {
		return result.Error
	}
	log.Println("get order data:")
	for _, orderStudent := range orderStudents {
		log.Printf("orderStudent=[%+v]", orderStudent)
	}

	//2.偏移量查询
	var limitStudents []mysql.Student
	if result := db.Order("age desc,id asc").Offset(2).Limit(2).Find(&limitStudents); result.Error != nil {
		return result.Error
	}
	log.Println("get offset and limit data:")
	for _, limitStudent := range limitStudents {
		log.Printf("limitStudent=[%+v]", limitStudent)
	}

	//3.group by and having查询
	var groupResult *groupResult
	if result := db.Select("age,count(1) as total").
		Model(&mysql.Student{}).
		Group("age").
		Having("total > 1").
		Scan(&groupResult); result.Error != nil {
		return result.Error
	}
	log.Println("get group by data:")
	log.Printf("result=[%+v]", groupResult)

	//4.join查询
	var teacher *mysql.Teacher
	db.Where("name = ?", "王老师").Find(&teacher)

	var joinStudents []mysql.Student
	db.Select("students.id,students.name,students.age,students.grade,students.class").
		Model(&mysql.Student{}).
		Joins("inner join teacher_and_students tas on students.id = tas.student_id").
		Where("tas.teacher_id = ?", teacher.ID).
		Order("students.id desc").
		Scan(&joinStudents)
	log.Println("get join data:")
	for _, joinStudent := range joinStudents {
		log.Printf("joinStudent=[%+v]", joinStudent)
	}

	//5.默认返回
	return nil
}
