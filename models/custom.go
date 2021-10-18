package models

import (
	"SIMS/global"
	"SIMS/internal/entity"
	"SIMS/utils/msg"
	"errors"
	"gorm.io/gorm"
)

type Custom struct {
	BaseModel
	CName    string `json:"c_name"    gorm:"comment:'客户名称'"`
	CNumber  string `json:"c_number"  gorm:"comment:'顾客编码'"`
	Phone    string `json:"phone"     gorm:"comment:'手机号码'"`
	Address  string `json:"address"   gorm:"comment:'收件件地址'"`
	Discount int    `json:"discount"  gorm:"comment:'折扣'"`
	Level    int    `json:"level"     gorm:"level:'客户等级'"`
}

func (c *Custom) CheckCustomExit() (err error, t Custom) {
	hasCustom := global.GDB.Where("c_name = ?", c.CName).First(&t).Error
	hasResult := errors.Is(hasCustom, gorm.ErrRecordNotFound)
	if !hasResult {
		return msg.DuplicatedData, t
	}
	return
}

func (c *Custom) CreateCustom() (err error) {
	err = entity.CheckExist(global.GDB, c, "c_name", c.CName)
	if err != nil {
		return err
	}
	err = global.GDB.Create(&c).Error
	if err != nil {
		return msg.CreatedFail
	}
	return nil
}
