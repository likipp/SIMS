package models

import (
	"SIMS/global"
	"SIMS/internal/entity"
	"SIMS/internal/schema"
	"SIMS/utils/msg"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type CustomLevel struct {
	BaseModel
	Name     string
	Discount int
	Customs  []Custom `gorm:"foreignKey:Level"`
}

type QueryParams struct {
	schema.PaginationParam
	Name     string `form:"name"`
	Discount int    `form:"discount"`
}

func GetCustomLevelDB(db *gorm.DB) *gorm.DB {
	return entity.GetDBWithModel(db, new(CustomLevel))
}

func (c *CustomLevel) CheckCustomLevelExist() (err error, t CustomLevel) {
	hasCustomLevel := global.GDB.Where("name = ?", c.Name).First(&t).Error
	hasResult := errors.Is(hasCustomLevel, gorm.ErrRecordNotFound)
	if !hasResult {
		return msg.DuplicatedData, t
	}
	return
}

func (c *CustomLevel) CreateCustomLevel() (err error) {
	err, _ = c.CheckCustomLevelExist()
	if err != nil {
		return err
	}
	err = global.GDB.Create(c).Error
	if err != nil {
		return
	}
	return err
}

func (c *CustomLevel) GetCustomLevel(params QueryParams) (err error, List []CustomLevel, total int64) {
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
	err = schema.QueryPaging(params.PaginationParam)
	if err != nil {
		return msg.GetFail, nil, 0
	}
	return nil, List, int64(len(List))
}

func (c *CustomLevel) UpdateCustomLevel(id int) (err error) {
	var t CustomLevel
	var i CustomLevel
	i.ID = id
	err, t = i.CheckCustomLevelExist()
	if err == nil {
		return msg.NotFound
	}
	err = global.GDB.Model(&t).Where("id = ?", id).Updates(&c).Error
	if err != nil && err == msg.DoNothing {
		return err
	}
	return nil
}

func (c *CustomLevel) DeleteCustomLevel() (err error) {
	err, _ = c.CheckCustomLevelExist()
	if err == nil {
		return msg.GetFail
	}
	err = global.GDB.Delete(&c).Error
	if err != nil {
		return msg.DeletedFail
	}
	return err
}
