package models

import (
	"SIMS/global"
	"SIMS/internal/entity"
	"SIMS/utils/msg"
	"gorm.io/gorm"
)

type WareHouse struct {
	BaseModel
	Name   string `json:"c_name"    gorm:"comment:'客户名称'"`
	Number string `json:"c_number"  gorm:"comment:'顾客编码'"`
}

func GetWareHouseDB() *gorm.DB {
	return entity.GetDBWithModel(global.GDB, new(WareHouse))
}

func (w *WareHouse) CreateWareHouse() (err error, success bool) {
	db := GetWareHouseDB()
	err = entity.CheckExist(db, "name", w.Name)
	if err != nil {
		return err, false
	}
	err = db.Create(w).Error
	if err != nil {
		return msg.CreatedFail, false
	}
	return err, true
}