package models

import (
	"SIMS/global"
	"SIMS/internal/entity"
	"SIMS/utils/msg"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

type Brand struct {
	BaseModel
	Name     string     `json:"name" gorm:"comment:'名称'"`
	Number   string     `json:"number" gorm:"comment:'编码'"`
	Products []Products `json:"products" gorm:"comment:'产品';foreignKey:brand"`
}

type BrandSelect struct {
	Value string `json:"value"`
	Label string `json:"label"`
	Key   string `json:"key"`
}

type BrandTree struct {
	Title string `json:"title"`
	Key   string    `json:"key"`
}

func GetBrandDB(db *gorm.DB) *gorm.DB {
	return entity.GetDBWithModel(db, new(Brand))
}

func (b *Brand) CreateBrand() (err error, success bool) {
	db := GetBrandDB(global.GDB)
	err = entity.CheckExist(db, "name", b.Name)
	if err != nil {
		return err, false
	}
	err = db.Create(b).Error
	if err != nil {
		return msg.CreatedFail, false
	}
	return msg.CreatedSuccess, true
}

func GetBrandSelectList(param string) (err error, list []BrandSelect, success bool) {
	var bs BrandSelect
	var bsl []BrandSelect
	var bl []Brand
	con := fmt.Sprintf("%s%s%s", "%", param, "%")
	db := GetBrandDB(global.GDB)
	if param != "" {
		err = db.Where("name like ?", con).Find(&bl).Error
		if err != nil {
			return msg.GetFail, list, false
		}
	}
	if err = db.Find(&bl).Error; err != nil {
		return msg.GetFail, list, false
	}
	for i, _ := range bl {
		bs.Value = strconv.Itoa(bl[i].ID)
		bs.Key = strconv.Itoa(bl[i].ID)
		bs.Label = bl[i].Name
		bsl = append(bsl, bs)
	}
	return msg.GetSuccess, bsl, true
}

func GetBrandTree() (error, []BrandTree, bool) {
	db := GetBrandDB(global.GDB)
	var bl []Brand
	var btl []BrandTree
	var bt BrandTree
	if err := db.Find(&bl).Error; err != nil {
		return msg.GetFail, nil, false
	}
	for _, brand := range bl {
		bt.Title = brand.Name
		bt.Key = brand.Number
		btl = append(btl, bt)
	}
	return msg.GetSuccess, btl, true
}
