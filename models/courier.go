package models

import (
	orm "SIMS/init/database"
	"SIMS/utils/msg"
	"errors"
	"gorm.io/gorm"
)

// 快递公司

type Courier struct {
	BaseModel
	CName   string `gorm:"column:c_name; comment:'快递名称'; not null; unique" json:"c_name"`
	CNumber string `gorm:"column:c_number; comment:'快递编号'" json:"c_number"`
	//Shipment
}

func (Courier) TableName() string {
	return "courier"
}

// CheckCourierExist 查询快递公司是否已经存在
func (c *Courier) CheckCourierExist() (err error, t Courier) {
	if c.CName != "" {
		hasCourier := orm.DB.Where("c_name = ?", c.CName).First(&t).Error
		hasResult := errors.Is(hasCourier, gorm.ErrRecordNotFound)
		if !hasResult {
			err = errors.New("快递公司已经存在,请检查已有数据")
			return err, t
		}
	}

	if c.ID > 0 {
		hasCourier := orm.DB.Where("id = ?", c.ID).First(&t).Error
		hasResult := errors.Is(hasCourier, gorm.ErrRecordNotFound)
		if !hasResult {
			err = errors.New("快递公司已经存在,请检查已有数据")
			return err, t
		}
	}
	return
}

// CreateCourier 创建快递公司
func (c *Courier) CreateCourier() (err error) {
	err, _ = c.CheckCourierExist()
	if err != nil {
		return err
	}
	err = orm.DB.Create(c).Error
	if err != nil {
		return
	}
	return err
}

// UpdateCourier 更新快递公司
func (c *Courier) UpdateCourier(id int) (err error) {
	var t Courier
	var i Courier
	i.ID = id
	err, t = i.CheckCourierExist()
	if err == nil {
		err = errors.New("快递公司不存在,请先登记")
		return err
	}
	err = orm.DB.Model(&t).Where("id = ?", id).Updates(&c).Error
	//m := fmt.Sprintf("%s", err) == msg.DoNothing
	if err != nil && err == msg.DoNothing {
		return err
	}
	return nil
}

// DeleteCourier 删除快递公司
func (c *Courier) DeleteCourier() (err error) {
	err, _ = c.CheckCourierExist()
	if err == nil {
		err = errors.New("快递公司不存在,请先登记")
		return err
	}
	err = orm.DB.Delete(&c).Error
	if err != nil {
		return err
	}
	return err
}

//func (c *Courier) BeforeCreate(tx *gorm.DB) (err error) {
//	if !c.IsValid() {
//
//	}
//}

func (c *Courier) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed("CName") {
		return msg.DoNothing
	}
	return nil
}
