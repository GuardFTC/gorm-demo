// Package mysql @Author:冯铁城 [17615007230@163.com] 2023-10-11 09:01:26
package mysql

// TeacherAndStudent 学生老师关联表模型
type TeacherAndStudent struct {
	ID        uint `gorm:"type:int(8);primaryKey;autoIncrement;comment:主键ID"`
	TeacherId uint `gorm:"type:int(8);comment:教师ID"`
	StudentId uint `gorm:"type:int(8);comment:学生ID"`
}
