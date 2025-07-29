// Package _gen @Author:冯铁城 [17615007230@163.com] 2025-07-23 16:51:50
package _gen

import (
	"gorm-demo/_gen/query"
	"log"
)

func Select(q *query.Query) error {

	//1.不携带查询条件
	err := selectWithoutParam(q)
	if err != nil {
		return err
	}

	//2.根据主键ID查询
	err = selectByPrimaryKey(q)
	if err != nil {
		return err
	}

	//3.条件查询
	err = selectByCondition(q)
	if err != nil {
		return err
	}

	//4.特殊查询
	err = specialSelect(q)
	if err != nil {
		return err
	}

	//默认返回
	return nil
}

// 不携带条件查询
func selectWithoutParam(q *query.Query) error {

	//1.全量查询
	if find, err := q.Student.Find(); err != nil {
		return err
	} else {
		log.Printf("get all data:")
		for _, student := range find {
			log.Printf("student=[%+v]", student)
		}
	}

	//2.查询根据主键ID正序排列的第一个数据
	if first, err := q.Student.First(); err != nil {
		return err
	} else {
		log.Printf("get order by id asc first data:")
		log.Printf("first=[%+v]", first)
	}

	//3.查询根据主键ID倒序排列的第一个数据
	if last, err := q.Student.Last(); err != nil {
		return err
	} else {
		log.Printf("get order by id desc first data:")
		log.Printf("last=[%+v]", last)
	}

	//4.任意获取一条数据
	if take, err := q.Student.Take(); err != nil {
		return err
	} else {
		log.Printf("get no order by data:")
		log.Printf("take=[%+v]", take)
	}

	//5.强制从主库读取，根据主键ID倒序排列的第一个数据
	if force, err := q.Student.WriteDB().Last(); err != nil {
		return err
	} else {
		log.Printf("get order by id desc first data from write db:")
		log.Printf("force=[%+v]", force)
	}

	//6.默认返回
	return nil
}

// 根据主键ID查询
func selectByPrimaryKey(q *query.Query) error {

	//1.根据主键ID查询单个数据
	if first, err := q.Teacher.Where(q.Teacher.ID.Eq(1)).First(); err != nil {
		return err
	} else {
		log.Printf("get by id data:")
		log.Printf("first=[%+v]", first)
	}

	//2.根据主键ID查询多个数据
	if find, err := q.Teacher.Where(q.Teacher.ID.In(1, 2)).Find(); err != nil {
		return err
	} else {
		log.Printf("get by id data:")
		for _, student := range find {
			log.Printf("student=[%+v]", student)
		}
	}

	//3.默认返回
	return nil
}

// 条件查询
func selectByCondition(q *query.Query) error {

	//1.等值查询
	if find, err := q.Student.Where(q.Student.Name.Eq("夏洛")).Find(); err != nil {
		return err
	} else {
		log.Printf("get where equal data:")
		for _, student := range find {
			log.Printf("student=[%+v]", student)
		}
	}

	//2.不等值查询
	if find, err := q.Student.Where(q.Student.Age.Neq(18)).Find(); err != nil {
		return err
	} else {
		log.Printf("get where not equal data:")
		for _, student := range find {
			log.Printf("student=[%+v]", student)
		}
	}

	//3.in查询
	if find, err := q.Student.Where(q.Student.Where(q.Student.Age.In(17, 18))).Find(); err != nil {
		return err
	} else {
		log.Printf("get where in data:")
		for _, student := range find {
			log.Printf("student=[%+v]", student)
		}
	}

	//4.like查询
	if find, err := q.Student.Where(q.Student.Email.Like("%zhang%")).Find(); err != nil {
		return err
	} else {
		log.Printf("get where like data:")
		for _, student := range find {
			log.Printf("student=[%+v]", student)
		}
	}

	//5.大于/大于等于查询
	if find, err := q.Student.Where(q.Student.Age.Gt(18)).Find(); err != nil {
		return err
	} else {
		log.Printf("get where gq/ge data:")
		for _, student := range find {
			log.Printf("student=[%+v]", student)
		}
	}

	//6.小于/小于等于查询
	if find, err := q.Student.Where(q.Student.Age.Lte(18)).Find(); err != nil {
		return err
	} else {
		log.Printf("get where lq/le data:")
		for _, student := range find {
			log.Printf("student=[%+v]", student)
		}
	}

	//7.AND查询
	if find, err := q.Student.Where(
		q.Student.Where(q.Student.Age.Eq(18)),
		q.Student.Where(q.Student.Email.Like("%y%")),
	).Find(); err != nil {
		return err
	} else {
		log.Printf("get where and data:")
		for _, student := range find {
			log.Printf("student=[%+v]", student)
		}
	}

	//8.OR查询
	if find, err := q.Student.Where(q.Student.Age.Eq(18)).Or(q.Student.Email.Like("%y%")).Find(); err != nil {
		return err
	} else {
		log.Printf("get where or data:")
		for _, student := range find {
			log.Printf("student=[%+v]", student)
		}
	}

	//9.NOT查询
	if find, err := q.Student.Not(q.Student.Age.In(18)).Find(); err != nil {
		return err
	} else {
		log.Printf("get where not data:")
		for _, student := range find {
			log.Printf("student=[%+v]", student)
		}
	}

	//10.查询指定字段
	if find, err := q.Student.Select(q.Student.ID, q.Student.Name).Where(q.Student.Age.Eq(18)).Find(); err != nil {
		return err
	} else {
		log.Printf("get where select data:")
		for _, student := range find {
			log.Printf("student=[%+v]", student)
		}
	}

	//11.查询数量
	if count, err := q.Student.Where(q.Student.Age.Eq(18)).Count(); err != nil {
		return err
	} else {
		log.Println("get count:")
		log.Printf("count=[%+v]", count)
	}

	//12.默认返回
	return nil
}

// ageGroup接收结果结构体
type ageGroup struct {
	Age   int32
	Count int32
}

// studentVo接收结果结构体
type studentVo struct {
	ID          int32  `gorm:"column:id;"`           // 主键ID
	Name        string `gorm:"column:name;"`         // 名称
	Email       string `gorm:"column:email;"`        // 邮箱
	Age         int32  `gorm:"column:age;"`          // 年龄
	Grade       int32  `gorm:"column:grade;"`        // 年级编号
	Class       int32  `gorm:"column:class;"`        // 班级编号
	TeacherName string `gorm:"column:teacher_name;"` // 老师名称
}

// 特殊查询
func specialSelect(q *query.Query) error {

	//1.排序查询
	if find, err := q.Student.Order(q.Student.Age.Desc(), q.Student.ID.Asc()).Find(); err != nil {
		return err
	} else {
		log.Printf("get group by data:")
		for _, group := range find {
			log.Printf("group=[%+v]", group)
		}
	}

	//2.偏移量查询
	if find, err := q.Student.Order(q.Student.Age.Desc(), q.Student.ID.Asc()).Offset(2).Limit(2).Find(); err != nil {
		return err
	} else {
		log.Println("get offset and limit data:")
		for _, limitStudent := range find {
			log.Printf("limitStudent=[%+v]", limitStudent)
		}
	}

	//3.Group By Having查询
	if rows, err := q.Student.
		Select(q.Student.Age, q.Student.Age.Count().As("count")).
		Group(q.Student.Age).
		Having(q.Student.Age.Count().Gt(1)).
		Rows(); err != nil {
		return err
	} else {
		log.Println("get group by data:")
		for rows.Next() {
			var groupResult ageGroup
			if err := rows.Scan(&groupResult.Age, &groupResult.Count); err != nil {
				return err
			}
			log.Printf("group=[%+v]", groupResult)
		}
	}

	//4.Join查询
	s := q.Student.As("s")
	tas := q.TeacherAndStudent.As("tas")
	if find, err := s.
		Select(s.ID, s.Name, s.Age, s.Grade, s.Class, s.Email, s.CreateTime, s.UpdateTime).
		Join(tas, s.ID.EqCol(tas.StudentID)).
		Where(tas.TeacherID.Eq(1)).
		Order(s.ID.Desc()).
		Find(); err != nil {
		return err
	} else {
		log.Println("get join data:")
		for _, joinStudent := range find {
			log.Printf("joinStudent=[%+v]", joinStudent)
		}
	}

	//5.Join查询-非DB对应结构体
	var studentVos []*studentVo
	s = q.Student.As("s")
	tas = q.TeacherAndStudent.As("tas")
	t := q.Teacher.As("t")
	if err := s.Select(s.ID, s.Name, s.Age, s.Grade, s.Class, s.Email, t.Name.As("teacher_name")).
		Join(tas, s.ID.EqCol(tas.StudentID)).
		Join(t, tas.TeacherID.EqCol(t.ID)).
		Where(t.ID.Eq(1)).
		Order(s.ID.Desc()).
		Scan(&studentVos); err != nil {
		return err
	} else {
		log.Println("get join data:")
		for _, studentVo := range studentVos {
			log.Printf("studentVo=[%+v]", studentVo)
		}
	}

	//6.默认返回
	return nil
}
