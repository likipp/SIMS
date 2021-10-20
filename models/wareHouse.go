package models

import (
	"SIMS/global"
	"SIMS/internal/entity"
	"SIMS/utils/msg"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gorm.io/gorm"
)

type WareHouse struct {
	BaseModel
	Name   string `json:"name"    gorm:"comment:'仓库名称'"`
	Number string `json:"number"  gorm:"comment:'仓库编码'"`
}

func (w *WareHouse) Validate() error {
	err := validation.ValidateStruct(w,
		validation.Field(&w.Name, validation.Required.Error("仓库名称不能为空")),
	)
	return err
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
	return msg.CreatedSuccess, true
}
