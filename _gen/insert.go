// @Author:冯铁城 [17615007230@163.com] 2025-07-23 15:18:10
package _gen

import (
	"gorm-demo/_gen/model"
	"gorm-demo/_gen/query"
)

// Insert 保存数据
func Insert(tx *query.Query) error {

	//1.单个数据保存
	student := &model.Student{Name: "夏洛", Email: "xialuo@xhs.com", Age: 18, Grade: 3, Class: 3}
	if err := tx.Student.Create(student); err != nil {
		return err
	}

	//2.批量数据保存
	students := make([]*model.Student, 0)
	students = append(students, &model.Student{Name: "马冬梅", Email: "madongmei@xltfn.cn", Age: 18, Grade: 3, Class: 1})
	students = append(students, &model.Student{Name: "大傻春", Email: "dashachun@xltfn.cn", Age: 21, Grade: 3, Class: 1})
	students = append(students, &model.Student{Name: "袁华", Email: "yuanhua@xltfn.cn", Age: 17, Grade: 4, Class: 2})
	students = append(students, &model.Student{Name: "秋雅", Email: "qiuya@xltfn.cn", Age: 18, Grade: 3, Class: 1})
	students = append(students, &model.Student{Name: "张扬", Email: "zhangyang@xltfn.cn", Age: 18, Grade: 3, Class: 1})
	students = append(students, &model.Student{Name: "陈凯", Email: "chenkai@xltfn.cn", Age: 21, Grade: 8, Class: 3})
	if err := tx.Student.Create(students...); err != nil {
		return err
	}

	//3.分批次批量保存
	teachers := make([]*model.Teacher, 0)
	teachers = append(teachers, &model.Teacher{Name: "王老师", Email: "wanglaoshi@xltfn.cn", Age: 43, Subject: "语文"})
	teachers = append(teachers, &model.Teacher{Name: "校长", Email: "xiaozhang@xltfn.cn", Age: 65, Subject: "学校"})
	if err := tx.Teacher.CreateInBatches(teachers, 10); err != nil {
		return err
	}

	//4.保存中间表数据
	teacherAndStudents := make([]*model.TeacherAndStudent, 0)
	teacherAndStudents = append(teacherAndStudents, &model.TeacherAndStudent{TeacherID: teachers[0].ID, StudentID: student.ID})
	for _, student := range students {
		teacherAndStudents = append(teacherAndStudents, &model.TeacherAndStudent{TeacherID: teachers[0].ID, StudentID: student.ID})
	}
	if err := tx.TeacherAndStudent.Create(teacherAndStudents...); err != nil {
		return err
	}

	//5.默认返回
	return nil
}
