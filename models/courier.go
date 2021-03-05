package models

import (
	orm "SIMS/init/database"
	"SIMS/utils/msg"
	"errors"
	"fmt"
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

	if c.CName != "" {
		hasCourier := orm.DB.Where("c_name = ?", c.CName).First(&t).Error
		hasResult := errors.Is(hasCourier, gorm.ErrRecordNotFound)
		if !hasResult {
			err = errors.New("供应商已经存在,请检查已有数据")
			return err
		}
	}

	if c.ID > 0 {
		hasCourier := orm.DB.Where("id = ?", c.ID).First(&t).Error
		hasResult := errors.Is(hasCourier, gorm.ErrRecordNotFound)
		if !hasResult {
			err = errors.New("供应商已经存在,请检查已有数据")
			return err
		}
	}
	return
}

// 创建供应商
func (c *Courier) CreateCourier() (err error) {
	err = c.CheckCourierExist()
	if err != nil {
		return err
	}
	err = orm.DB.Create(c).Error
	if err != nil {
		return
	}
	return err
}

func (c *Courier) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed("c_name") {
		return errors.New(msg.DoNothing)
	}
	return nil
}

// 更新供应商
func (c *Courier) UpdateCourier(id int) (err error) {
	var t Courier
	var i Courier
	fmt.Println(c, "Courier")
	i.ID = id
	err = i.CheckCourierExist()
	if err == nil {
		err = errors.New("供应商不存在,请先登记")
		return err
	}
	err = orm.DB.Model(&t).Where("id = ?", id).Updates(c).Error
	m := fmt.Sprintf("%s", err) != msg.DoNothing
	if err != nil && m {
		return err
	}
	return nil
}

// 删除供应商
func (c *Courier) DeleteCourier() (err error) {
	err = c.CheckCourierExist()
	if err == nil {
		err = errors.New("供应商不存在,请先登记")
		return err
	}
	err = orm.DB.Delete(&c).Error
	if err != nil {
		return err
	}
	return err
}
