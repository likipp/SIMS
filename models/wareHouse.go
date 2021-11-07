package models

import (
	"SIMS/global"
	"SIMS/internal/entity"
	"SIMS/utils/msg"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gorm.io/gorm"
)

type WareHouse struct {
	BaseModel
	Name   string `json:"name"    gorm:"comment:'仓库名称'"`
	Number string `json:"number"  gorm:"comment:'仓库编码'"`
}

type WareHouseSelect struct {
	Value string `json:"value"`
	Label string `json:"label"`
	Key   string `json:"key"`
}

func (w *WareHouse) Validate() error {
	err := validation.ValidateStruct(w,
		validation.Field(&w.Name, validation.Required.Error("仓库名称不能为空")),
	)
	return err
}

func GetWareHouseDB() *gorm.DB {
	return entity.GetDBWithModel(global.GDB, new(WareHouse))
}

func (w *WareHouse) CreateWareHouse() (err error, success bool) {
	db := GetWareHouseDB()
	err = entity.CheckExist(db, "name", w.Name)
	if err != nil {
		return err, false
	}
	err = db.Create(w).Error
	if err != nil {
		return msg.CreatedFail, false
	}
	return msg.CreatedSuccess, true
}

//
//func GetWareHouseList(param string) (err error, list []WareHouseSelect, success bool) {
//	var w WareHouseSelect
//	var ws []WareHouseSelect
//	var wsl []WareHouse
//	con := fmt.Sprintf("%s%s%s", "%", param, "%")
//	db := GetWareHouseDB()
//	if param != "" {
//		err = db.Where("name like ? or number like ?", con, con).Find(&wsl).Error
//		if err != nil {
//			return msg.GetFail, list, false
//		}
//	}
//	if err = db.Find(&wsl).Error; err != nil {
//		return msg.GetFail, list, false
//	}
//	for _, pro := range wsl {
//		w.Value = pro.Number
//		w.Label = pro.Name
//		w.Key = pro.Number
//		ws = append(ws, w)
//	}
//	return msg.GetSuccess, ws, true
//}

func GetWareHouseSelectList(param string) (err error, list []WareHouseSelect, success bool) {
	//var p ProductsSelect
	var ps []WareHouseSelect
	//var psl []Products
	con := fmt.Sprintf("%s%s%s", "%", param, "%")
	var selectData = "name as label, number as value"
	db := GetWareHouseDB()
	if param != "" {
		err = db.Select(selectData).Where("name like ? or number like ?", con, con).Find(&ps).Error
		if err != nil {
			return msg.GetFail, list, false
		}
	}
	if err = db.Select(selectData).Find(&ps).Error; err != nil {
		return msg.GetFail, list, false
	}
	for i, _ := range ps {
		ps[i].Key = ps[i].Value
	}
	return msg.GetSuccess, ps, true
}
