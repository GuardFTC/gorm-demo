// Package common @Author:冯铁城 [17615007230@163.com] 2025-07-18 11:10:30
package common

import (
	"log"

	"gorm.io/gorm"
)

func CheckTransactionError(db *gorm.DB, fc func(tx *gorm.DB) error) {
	CheckError(db.Transaction(fc))
}

// CheckError 错误检查
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
