package models

import (
	"SIMS/utils/msg"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type BillEntry struct {
	BaseModel
	HeaderID  int     `json:"header_id" gorm:"comment:'表头ID'"`
	PNumber   string  `json:"p_number" gorm:"comment:'产品代码'"`
	PName     string  `json:"p_name" gorm:"comment:'产品名称'"`
	WareHouse int     `json:"ware_house" gorm:"comment:'仓库'"`
	ExQTY     int     `json:"ex_qty" gorm:"comment:'出库数量'"`
	InQTY     int     `json:"in_qty" gorm:"comment:'入库数量'"`
	UnitPrice float32 `json:"unit_price" gorm:"comment:'单价'"`
	Discount  float32 `json:"discount" gorm:"comment:'折扣'"`
	Total     float32 `json:"total" gorm:"comment:'金额'"`
}

func (be *BillEntry) Validate() error {
	err := validation.ValidateStruct(be,
		validation.Field(&be.PNumber, validation.Required.Error("产品代码不能为空")),
		validation.Field(&be.PName, validation.Required.Error("产品名称不能为空")),
		validation.Field(&be.WareHouse, validation.Required.Error("仓库不能为空")),
		validation.Field(&be.UnitPrice, validation.Required.Error("单价不能为空")),
		validation.Field(&be.Total, validation.Required.Error("金额不能为空")),
		//validation.Field(&sb.InQTY, validation.When(s.StockType == global.In, validation.Required.Error("数量不能为空"))),
		//validation.Field(&sb.ExQTY, validation.When(s.StockType != global.In, validation.Required.Error("数量不能为空"))),
	)
	return err
}

func (be *BillEntry) InStockLog(tx *gorm.DB) (err error, success bool) {
	stock := GetWareHouseQtyWithProduct(be.WareHouse, be.PNumber)
	if stock.QTY > 0 {
		stock.QTY = be.InQTY + stock.QTY
		err, success = stock.UpdateStockWithTransaction(tx)
		if !success {
			return err, false
		}
		if err = tx.Create(&be).Error; err != nil {
			return msg.CreatedFail, false
		}
		return err, true
	}
	err = copier.Copy(&stock, &be)
	if err != nil {
		return msg.Copier, false
	}
	stock.QTY = be.InQTY + stock.QTY
	err, success = stock.CreateStockWithTransaction(tx)
	if !success {
		return err, false
	}
	if err = tx.Create(&be).Error; err != nil {
		return msg.CreatedFail, false
	}
	return msg.CreatedSuccess, true
}

func (be *BillEntry) ExStockLog(tx *gorm.DB) (err error, success bool) {
	stock := GetWareHouseQtyWithProduct(be.WareHouse, be.PNumber)
	if stock.QTY > 0 {
		if err, success = stock.CheckStock(be.ExQTY); !success {
			return err, false
		}
		stock.QTY = stock.QTY - be.ExQTY
		err, success = stock.UpdateStockWithTransaction(tx)
		if !success {
			return err, false
		}
		if err = tx.Create(&be).Error; err != nil {
			return msg.CreatedFail, false
		}
		return msg.CreatedSuccess, true
	}
	return msg.ExGTStock, false
}

//func GetStockIn() {
//	global.GDB.Where("stock_type = ?", global.In).Find()
//}
