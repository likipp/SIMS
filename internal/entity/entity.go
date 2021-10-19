package entity

import (
	"SIMS/utils/msg"
	"fmt"
	"gorm.io/gorm"
)

func GetDBWithModel(db *gorm.DB, i interface{}) *gorm.DB {
	return db.Model(i)
}

// CheckExist 检查是否有重复数据
func CheckExist(db *gorm.DB, param, field string) (err error) {
	var count int64
	query := fmt.Sprintf("%s = ?", param)
	db.Where(query, field).Count(&count)
	if count > 0 {
		return msg.DuplicatedData
	}
	return nil
}
