// Package common @Author:冯铁城 [17615007230@163.com] 2025-07-18 11:10:30
package common

import (
	"gorm-demo/_gen/query"
	"log"

	"gorm.io/gorm"
)

// CheckGormTransactionError 事务错误检查
func CheckGormTransactionError(db *gorm.DB, fc func(tx *gorm.DB) error) {
	CheckError(db.Transaction(fc))
}

// CheckGenTransactionError 事务错误检查
func CheckGenTransactionError(query *query.Query, fc func(tx *query.Query) error) {
	CheckError(query.Transaction(fc))
}

// CheckError 错误检查
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
