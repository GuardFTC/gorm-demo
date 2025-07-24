// Package _gen @Author:冯铁城 [17615007230@163.com] 2025-07-24 14:15:52
package _gen

import (
	"gorm-demo/_gen/query"
	"log"
)

// Delete 删除数据
func Delete(tx *query.Query) error {

	//1.根据主键ID删除单条数据
	if info, err := tx.Teacher.Where(tx.Teacher.ID.Eq(1)).Delete(); err != nil {
		return err
	} else {
		log.Printf("删除数据成功，受影响的行数：%d", info.RowsAffected)
	}

	//2.根据主键ID删除多条数据
	if info, err := tx.Student.Where(tx.Student.ID.In(1, 2, 3, 4, 5, 6, 7)).Delete(); err != nil {
		return err
	} else {
		log.Printf("删除数据成功，受影响的行数：%d", info.RowsAffected)
	}

	//3.根据条件删除数据
	if info, err := tx.TeacherAndStudent.Where(tx.TeacherAndStudent.ID.Gte(1)).Delete(); err != nil {
		return err
	} else {
		log.Printf("删除数据成功，受影响的行数：%d", info.RowsAffected)
	}

	//4.默认返回
	return nil
}
