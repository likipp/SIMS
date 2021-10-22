package models

import (
	"SIMS/global"
	"SIMS/internal/entity"
	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type Stock struct {
	BaseModel
	PName          string `json:"p_name"`
	PNumber        string    `json:"p_number"`
	WareHouse      int    `json:"ware_house"`
	QTY            int    `json:"qty"`
}

type StockWithAction struct {
	Stock
	Action string `json:"action"`
}

func (s *Stock) Validate() error {
	err := validation.ValidateStruct(s,
		validation.Field(&s.PNumber, validation.Required.Error("产品代码不能为空")),
		validation.Field(&s.PName, validation.Required.Error("产品名称不能为空")),
		validation.Field(&s.WareHouse, validation.Required.Error("仓库不能为空")),
		validation.Field(&s.QTY, validation.Required.Error("数量不能为空")),
		)
	return err
}

func GetStockDB() *gorm.DB {
	return entity.GetDBWithModel(global.GDB, new(Stock))
}

func GetWareHouseQtyWithProduct(wareHouse int , product string) (*Stock, bool) {
	var stock Stock
	var count int64
	db := GetStockDB()
	db.Where("ware_house = ? and p_number = ?", wareHouse, product).Find(&stock).Count(&count)
	if  count > 0 {
		return &stock, true
	}
	return &stock, false
}

//func (s *Stock) InStock() (err error, success bool) {
//	db := GetStockDB()
//	err = db.Create(s).Error
//	if err != nil {
//		return msg.CreatedFail, false
//	}
//	return msg.CreatedSuccess, true
//}
//
//func (s *Stock) ExStock() (err error, success bool) {
//	db := GetStockDB()
//	err = db.Create(s).Error
//	if err != nil && err == msg.ExGTStock {
//		return err, false
//	}
//	if err != nil && err != msg.ExGTStock {
//		return msg.CreatedFail, false
//	}
//	return msg.CreatedSuccess, true
//}

//func (s *Stock) BeforeCreate(tx *gorm.DB) (err error) {
//	var stock Stock
//	ex := s.QTY
//	wareHouse := s.WareHouse
//	pNumber := s.PNumber
//	err = tx.Select("qty").Where("ware_house = ? and p_number = ?", wareHouse, pNumber).Last(&stock).Error
//	if ex > stock.QTY {
//		err = msg.ExGTStock
//	}
//	return
//}