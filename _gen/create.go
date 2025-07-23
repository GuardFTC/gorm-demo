// @Author:冯铁城 [17615007230@163.com] 2025-07-23 15:18:10
package _gen

import (
	"gorm-demo/_gen/model"
	"gorm-demo/_gen/query"
)

// Insert 保存数据
func Insert(tx *query.Query) error {

	//1.单个数据保存
	student := &model.Student{Name: "夏洛", Email: "xialuo@xhs.com", Age: 18, Grade: 1, Class: 1}
	if err := tx.Student.Create(student); err != nil {
		return err
	}

	//默认返回
	return nil
}
