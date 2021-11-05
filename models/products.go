package models

import (
	"SIMS/global"
	"SIMS/internal/entity"
	"SIMS/internal/schema"
	"SIMS/utils/msg"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type Products struct {
	BaseModel
	PName     string  `json:"p_name" gorm:"comment:'产品名称'"`
	PNumber   string  `gorm:"index not null; comment:'产品编号'" json:"p_number"`
	Spec      string  `json:"p_spec" gorm:"comment:'规格型号'"`
	Price     float64 `json:"p_price" gorm:"comment:'单价'"`
	Brand     int     `json:"brand" gorm:"comment:'品牌'"`
	Unit      int     `json:"unit" gorm:"comment:'单位'"`
	WareHouse int     `json:"ware_house" gorm:"comment:'默认仓库'"`
	Mark      string  `json:"mark" gorm:"comment:'备注'"`
	Picture   string  `json:"picture" gorm:"default:'/favicon.ico'; comment:'图片'"`
}

//type ProductOrder struct {
//	Sorter string `form:"Sorter"`
//	//PNumberOrder string `form:"p_number_order"`
//	//BrandOrder string  `form:"brand_order"`
//	//IDOrder    string `form:"id_order"`
//}

type ProductQueryParams struct {
	schema.PaginationParam
	//ProductOrder
	Sorter  string `form:"sorter"`
	PName   string `form:"p_name"`
	PNumber string `form:"p_number"`
	Brand   int    `form:"brand"`
}

type ProductQueryResult struct {
	BaseModel
	PName   string  `json:"p_name"`
	PNumber string  `json:"p_number"`
	Spec    string  `json:"p_spec"`
	Price   float64 `json:"price"`
	Brand   int     `json:"brand"`
	Unit    int     `json:"unit"`
	Mark    string  `json:"mark"`
	Picture string
}

type ProductsSelect struct {
	Value string `json:"value"`
	Label string `json:"label"`
	Key   string `json:"key"`
}

func (p *Products) Validate() error {
	err := validation.ValidateStruct(p,
		validation.Field(&p.PName, validation.Required.Error("产品名称不能为空")),
		validation.Field(&p.Brand, validation.Required.Error("品牌不能为空")),
		validation.Field(&p.Unit, validation.Required.Error("单位不能为空")),
		validation.Field(&p.Price, validation.Min(1).Error("单价必须大于0"), validation.Required.Error("单价不能为空")),
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

func GetProductsSelectList(param string) (err error, list []ProductsSelect, success bool) {
	var p ProductsSelect
	var ps []ProductsSelect
	var psl []Products
	con := fmt.Sprintf("%s%s%s", "%", param, "%")
	db := GetProductsDB()
	if param != "" {
		err = db.Where("p_name like ? or p_number like ?", con, con).Find(&psl).Error
		if err != nil {
			return msg.GetFail, list, false
		}
	}
	if err = db.Find(&psl).Error; err != nil {
		return msg.GetFail, list, false
	}
	for _, pro := range psl {
		p.Value = pro.PName
		p.Label = pro.PNumber
		p.Key = pro.PNumber
		ps = append(ps, p)
	}
	return msg.GetSuccess, ps, true
}

func (p *Products) GetProductsList(params ProductQueryParams) (err error, List []ProductQueryResult, total int64, success bool) {
	db := GetProductsDB()
	var product Products
	//if v := params.PName; v != "" {
	//	db = db.Where("p_name like ?", fmt.Sprintf("%s%s%s", "%", v, "%")).Count(&total)
	//}
	//if v := params.PNumber; v != "" {
	//	db = db.Where("p_number like ?", fmt.Sprintf("%s%s%s", "%", v, "%")).Count(&total)
	//}
	//if v := params.Brand; v > 0 {
	//	db = db.Where("brand = ?", v).Count(&total)
	//}
	//if params.Brand <= 0 && params.PName == "" && params.PNumber == "" {
	//	db = db.Count(&total)
	//}
	err = copier.Copy(&product, &params)
	if err != nil {
		return msg.Copier, nil, 0, false
	}
	//order := utils.StructConvMap(params.ProductOrder)
	db.Scopes(schema.QueryPaging(params.PaginationParam), schema.QueryOrder(params.Sorter)).Find(&List)
	if err != nil {
		return msg.GetFail, nil, 0, false
	}
	return msg.GetSuccess, List, int64(len(List)), true
}

func (p *Products) UpdateProduct() (err error, success bool) {
	db := GetProductsDB()
	err = db.Where("id = ?", p.ID).Updates(&p).Error
	if err != nil {
		return msg.UpdatedFail, false
	}
	return msg.UpdatedSuccess, true
}