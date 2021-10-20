package models

type WareHouse struct {
	BaseModel
	Name   string `json:"c_name"    gorm:"comment:'客户名称'"`
	Number string `json:"c_number"  gorm:"comment:'顾客编码'"`
}
