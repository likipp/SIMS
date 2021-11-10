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

type CustomWithLevel struct {
	BaseModel
	CName     string `json:"c_name"`
	CNumber   string `json:"c_number"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Discount  int    `json:"discount"`
	LevelID   int    `json:"level_id"`
	LevelName string `json:"level_name"`
	Mark      string `json:"mark"`
}

type CustomQueryParams struct {
	schema.PaginationParam
	Name string `form:"name"`
}

type CustomSelect struct {
	Value     string `json:"value"`
	Label     string `json:"label"`
	ID        int    `json:"id"`
	Discount  int    `json:"discount"`
	LevelName string `json:"l_name"`
}

func GetCustomDB() *gorm.DB {
	return entity.GetDBWithModel(global.GDB, new(Custom))
}

func (c *Custom) CreateCustom() (err error, success bool) {
	db := GetCustomDB()
	err = entity.CheckExist(db, "c_name", c.CName)
	if err != nil {
		return err, false
	}
	err = db.Create(&c).Error
	if err != nil {
		return msg.CreatedFail, false
	}
	return msg.CreatedSuccess, true
}

func (c *Custom) GetList(params CustomQueryParams) (success bool, err error, List []CustomWithLevel, total int64) {
	var selectData = "customs.id, customs.c_name, customs.c_number, customs.phone, customs.address, customs.discount, customs.created_at, customs.create_by, customs.mark, custom_levels.discount, custom_levels.id as level_id,  custom_levels.name as level_name"
	var joinData = "join custom_levels on customs.level = custom_levels.id"
	db := GetCustomDB()

	if v := params.Name; v != "" {
		db = db.Select(selectData).Joins(joinData).Where("name like ?", fmt.Sprintf("%s%s%s", "%", v, "%")).Find(&List)
	} else {
		db.Select(selectData).Joins(joinData).Find(&List)

	}
	//err = schema.QueryPaging(params.PaginationParam)
	if err != nil {
		return false, msg.GetFail, nil, 0
	}
	return true, msg.GetSuccess, List, int64(len(List))
}

func GetCustomByID(id int) (success bool, err error, custom Custom) {
	db := GetCustomDB()
	if err = db.Where("id = ?", id).Find(&custom).Error; err != nil {
		return false, msg.GetFail, custom
	}
	return true, msg.GetSuccess, custom
}

func GetCustomsListWithQuery(param string) (err error, list []CustomSelect, success bool) {
	//var c CustomSelect
	var cs []CustomSelect
	//var cl []Custom
	con := fmt.Sprintf("%s%s%s", "%", param, "%")
	var selectData = "customs.id, customs.c_number as label, customs.c_name as value, custom_levels.name as l_name, custom_levels.discount"
	var joinData = "join custom_levels on customs.level = custom_levels.id"
	db := GetCustomDB()
	if param != "" {
		err = db.Select(selectData).Joins(joinData).Where("customs.c_name like ? or customs.c_number like ?", con, con).Find(&cs).Error
		if err != nil {
			return msg.GetFail, list, false
		}
	}
	if err = db.Select(selectData).Joins(joinData).Find(&cs).Error; err != nil {
		return msg.GetFail, list, false
	}
	return msg.GetSuccess, cs, true
}
