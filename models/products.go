package models

import (
	"SIMS/internal/entity"
	"gorm.io/gorm"
)

type Products struct {
	BaseModel
	PName     string      `json:"p_name" gorm:"comment:'产品名称'"`
	PNumber   string      `gorm:"index not null; comment:'产品编号'" json:"p_number"`
	Spec      string      `json:"p_spec" gorm:"comment:'规格型号'"`
	Price     string      `json:"p_price" gorm:"comment:'单价'"`
	Brand     int         `json:"brand" gorm:"comment:'品牌'"`
	WareHouse []WareHouse `json:"ware_house" gorm:"comment:'仓库';many2many:products_warehouse;foreignKey:Refer"`
	Refer     int         `gorm:"index:,unique"`
	Mark      string      `gorm:"comment:'备注'"`
	Picture   string      `json:"picture" gorm:"default:'/favicon.ico'; comment:'图片'"`
}

func GetProductsDB(db *gorm.DB) *gorm.DB {
	return entity.GetDBWithModel(db, new(Products))
}

func (c *Products) CreateProducts() (err error) {
	//
	return nil
}
