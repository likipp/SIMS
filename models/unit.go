package models

import (
	"SIMS/global"
	"SIMS/internal/entity"
	"SIMS/utils/msg"
	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type Unit struct {
	BaseModel
	Name string `json:"name" gorm:"comment:'名称'"`
}

func (u *Unit) Validate() error {
	err := validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Required.Error("单位名称不能为空")),
	)
	return err
}

func getUnitDB() *gorm.DB {
	return entity.GetDBWithModel(global.GDB, new(Unit))
}

func (u *Unit) CreateUnit() error {
	db := getUnitDB()
	err := entity.CheckExist(db, "name", u.Name)
	if err != nil {
		return err
	}
	err = db.Create(&u).Error
	if err != nil {
		return msg.CreatedFail
	}
	return msg.CreatedSuccess
}
