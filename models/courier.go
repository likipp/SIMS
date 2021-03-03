package models

import (
	orm "SIMS/init/database"
	"errors"
	"gorm.io/gorm"
)

// 快递公司

type Courier struct {
	BaseModel
	CName   string `gorm:"column:c_name" json:"c_name"`
	CNumber string `gorm:"column:c_number" json:"c_number"`
	//Shipment
}

func (Courier) TableName() string {
	return "courier"
}

// 查询供应商是否已经存在
func (c *Courier) CheckCourierExist() (err error) {
	var t Courier
	hasCourier := orm.DB.Where("c_name = ?", c.CName).First(&t).Error
	hasResult := errors.Is(hasCourier, gorm.ErrRecordNotFound)
	if !hasResult {
		err = errors.New("供应商已经存在,请检查已有数据")
		return
	}
	return
}


// 创建供应商
func(c *Courier) CreateCourier() (err error) {
	err = c.CheckCourierExist()
	if err != nil {
		return err
	}
	err = orm.DB.Create(c).Error
	if err != nil {
		return err
	}
	return
}


