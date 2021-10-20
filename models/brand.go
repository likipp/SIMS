package models

import (
	"SIMS/global"
	"SIMS/internal/entity"
	"SIMS/utils/msg"
	"gorm.io/gorm"
)

type Brand struct {
	BaseModel
	Name     string     `json:"name" gorm:"comment:'名称'"`
	Number   string     `json:"number" gorm:"comment:'编码'"`
	Products []Products `json:"products" gorm:"comment:'产品';foreignKey:brand"`
}

func GetBrandDB(db *gorm.DB) *gorm.DB {
	return entity.GetDBWithModel(db, new(Brand))
}

func (b *Brand) CreateBrand() (err error, success bool) {
	db := GetBrandDB(global.GDB)
	err = entity.CheckExist(db, "name", b.Name)
	if err != nil {
		return err, false
	}
	err = db.Create(b).Error
	if err != nil {
		return msg.CreatedFail, false
	}
	return msg.CreatedSuccess, true
}
