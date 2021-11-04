package models

import (
	"SIMS/global"
	"SIMS/internal/entity"
	"SIMS/internal/schema"
	"SIMS/utils/msg"
	"fmt"
	"gorm.io/gorm"
)

type CustomLevel struct {
	BaseModel
	Name     string   `json:"name"`
	Discount int      `json:"discount"`
	Customs  []Custom `gorm:"foreignKey:Level" json:"customs"`
}

type CustomLevelQueryParams struct {
	schema.PaginationParam
	Name     string `form:"name"`
	Discount int    `form:"discount"`
}

func GetCustomLevelDB(db *gorm.DB) *gorm.DB {
	return entity.GetDBWithModel(db, new(CustomLevel))
}

func (c *CustomLevel) CreateCustomLevel() (err error) {
	db := GetCustomLevelDB(global.GDB)
	err = entity.CheckExist(db, "name", c.Name)
	if err != nil {
		return err
	}
	err = db.Create(c).Error
	if err != nil {
		return msg.CreatedFail
	}
	return err
}

func (c *CustomLevel) GetList(params CustomLevelQueryParams) (success bool, err error, List []CustomLevel, total int64) {
	db := GetCustomLevelDB(global.GDB)
	if v := params.Name; v != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%s%s%s", "%", v, "%")).Find(&List)
	}
	if v := params.Discount; v > 0 {
		db = db.Where("discount = ?", v).Find(&List)
	}
	if params.Discount <= 0 && params.Name == "" {
		db = db.Find(&List)
	}
	//err = schema.QueryPaging(params.PaginationParam)
	if err != nil {
		return false, msg.GetFail, nil, 0
	}
	return true, msg.GetSuccess, List, int64(len(List))
}

func (c *CustomLevel) UpdateCustomLevel(id int) (err error) {
	var i CustomLevel
	i.ID = id
	db := GetCustomLevelDB(global.GDB)
	err = entity.CheckExist(db, "name", c.Name)
	if err == nil {
		return msg.NotFound
	}
	err = db.Where("id = ?", id).Updates(&c).Error
	if err != nil && err == msg.DoNothing {
		return err
	}
	return nil
}

func (c *CustomLevel) DeleteCustomLevel() (err error) {
	db := GetCustomLevelDB(global.GDB)
	err = entity.CheckExist(db, "name", c.Name)
	if err == nil {
		return msg.GetFail
	}
	err = db.Delete(&c).Error
	if err != nil {
		return msg.DeletedFail
	}
	return err
}
