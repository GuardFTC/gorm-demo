// Package mysql @Author:冯铁城 [17615007230@163.com] 2025-07-17 16:55:49
package mysql

import (
	"log"

	mysql_model "gorm-demo/mysql/model"
	"gorm.io/gorm"
)

// CreateTable 创建表
func CreateTable(db *gorm.DB) error {

	//1.创建学生表
	err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='学生表'").AutoMigrate(&mysql_model.Student{})
	if err != nil {
		return err
	}

	//2.创建老师表
	err = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='老师表'").AutoMigrate(&mysql_model.Teacher{})
	if err != nil {
		return err
	}

	//3.创建老师学生中间表
	err = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='老师学生表'").AutoMigrate(&mysql_model.TeacherAndStudent{})
	if err != nil {
		return err
	}

	//4.默认返回
	log.Printf("table create success")
	return nil
}

// DropTable 删除表
func DropTable(db *gorm.DB) error {

	//1.删除学生表
	if err := db.Migrator().DropTable(&mysql_model.Student{}); err != nil {
		return err
	}

	//2.删除老师表
	if err := db.Migrator().DropTable(&mysql_model.Teacher{}); err != nil {
		return err
	}

	//3.删除老师学生中间表
	if err := db.Migrator().DropTable(&mysql_model.TeacherAndStudent{}); err != nil {
		return err
	}

	//4.默认返回
	log.Printf("table drop success")
	return nil
}
