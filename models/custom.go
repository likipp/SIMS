package models

import (
	"SIMS/global"
	"SIMS/internal/entity"
	"SIMS/internal/schema"
	"SIMS/utils/msg"
	"fmt"
	"gorm.io/gorm"
)

type Custom struct {
	BaseModel
	CName    string `json:"c_name"    gorm:"comment:'客户名称'"`
	CNumber  string `json:"c_number"  gorm:"comment:'顾客编码'"`
	Phone    string `json:"phone"     gorm:"comment:'手机号码'"`
	Address  string `json:"address"   gorm:"comment:'收件件地址'"`
	Discount int    `json:"discount"  gorm:"comment:'折扣'"`
	Level    int    `json:"level"     gorm:"comment:'客户等级'"`
	Mark     string `json:"mark"      gorm:"comment:'备注'"`
}

type CustomQueryParams struct {
	schema.PaginationParam
	Name string `form:"name"`
}

func GetCustomDB() *gorm.DB {
	return entity.GetDBWithModel(global.GDB, new(Custom))
}

func (c *Custom) CreateCustom() (err error) {
	db := GetCustomDB()
	err = entity.CheckExist(db, "c_name", c.CName)
	if err != nil {
		return err
	}
	err = db.Create(&c).Error
	if err != nil {
		return msg.CreatedFail
	}
	return nil
}

func (c *Custom) GetList(params CustomQueryParams) (success bool, err error, List []Custom, total int64) {
	db := GetCustomDB()
	if v := params.Name; v != "" {
		db = db.Preload("CustomLevel").Where("name like ?", fmt.Sprintf("%s%s%s", "%", v, "%")).Find(&List)
	} else {
		db = db.Preload("Level").Find(&List)
	}
	err = schema.QueryPaging(params.PaginationParam)
	if err != nil {
		return false, msg.GetFail, nil, 0
	}
	return true, msg.GetSuccess, List, int64(len(List))
}
