package models

import (
	"SIMS/global"
	"SIMS/internal/entity"
	"SIMS/utils/msg"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type Products struct {
	BaseModel
	PName   string `json:"p_name" gorm:"comment:'产品名称'"`
	PNumber string `gorm:"index not null; comment:'产品编号'" json:"p_number"`
	Spec    string `json:"p_spec" gorm:"comment:'规格型号'"`
	Price   int    `json:"p_price" gorm:"comment:'单价'"`
	Brand   int    `json:"brand" gorm:"comment:'品牌'"`
	//WareHouse []WareHouse `json:"ware_house" gorm:"comment:'仓库';many2many:products_warehouse;foreignKey:Refer"`
	//Refer     int         `gorm:"index:,unique"`
	WareHouse int    `json:"ware_house" gorm:"comment:'默认仓库'"`
	Mark      string `gorm:"comment:'备注'"`
	Picture   string `json:"picture" gorm:"default:'/favicon.ico'; comment:'图片'"`
}

type ProductsSelect struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

func (p *Products) Validate() error {
	//validation.Min(1), validation.In(is.Float, is.Int),
	err := validation.ValidateStruct(p,
		validation.Field(&p.PName, validation.Required.Error("产品名称不能为空")),
		validation.Field(&p.Brand, validation.Required.Error("品牌不能为空")),
		validation.Field(&p.Price, validation.Min(1).Error("单价必须大于1"), validation.Required.Error("单价不能为空")),
	)
	return err
}

func GetProductsDB() *gorm.DB {
	return entity.GetDBWithModel(global.GDB, new(Products))
}

func (p *Products) CreateProducts() (err error, success bool) {
	db := GetProductsDB()
	err = entity.CheckExist(db, "p_name", p.PName)
	err = db.Create(p).Error
	if err != nil {
		return msg.CreatedFail, false
	}
	return msg.CreatedSuccess, true
}

func GetProductsList() (err error, list []ProductsSelect, success bool) {
	var p ProductsSelect
	var ps []ProductsSelect
	var psl []Products
	db := GetProductsDB()
	if err = db.Find(&psl).Error; err != nil {
		return msg.GetFail, list, false
	}
	for _, pro := range psl {
		p.Value = pro.PNumber
		p.Label = pro.PName
		ps = append(ps, p)
	}
	fmt.Println(ps, "ps信息")
	return msg.GetSuccess, ps, true
}
