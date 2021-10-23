package models

import (
	"SIMS/global"
	"SIMS/utils/msg"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/copier"
)

type StockBody struct {
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

func (sb *StockBody) Validate() error {
	err := validation.ValidateStruct(sb,
		validation.Field(&sb.PNumber, validation.Required.Error("产品代码不能为空")),
		validation.Field(&sb.PName, validation.Required.Error("产品名称不能为空")),
		validation.Field(&sb.WareHouse, validation.Required.Error("仓库不能为空")),
		validation.Field(&sb.UnitPrice, validation.Required.Error("单价不能为空")),
		validation.Field(&sb.Total, validation.Required.Error("金额不能为空")),
		//validation.Field(&sb.InQTY, validation.When(s.StockType == global.In, validation.Required.Error("数量不能为空"))),
		//validation.Field(&sb.ExQTY, validation.When(s.StockType != global.In, validation.Required.Error("数量不能为空"))),
	)
	return err
}

func (sb *StockBody) StockBodyLog(sh StockHeader) (err error, success bool) {
	stock, ok := GetWareHouseQtyWithProduct(sb.WareHouse, sb.PNumber)
	if ok {
		if sh.StockType == global.Ex {
			if stock.QTY < sb.ExQTY {
				return msg.ExGTStock, false
			}
			stock.QTY = stock.QTY - sb.ExQTY
			err = global.GDB.Model(stock).Update("qty", stock.QTY).Error
			if err != nil {
				return msg.UpdatedFail, false
			}
			return sb.StockLog()
		}
		stock.QTY = sb.InQTY + stock.QTY
		err = global.GDB.Model(stock).Update("qty", stock.QTY).Error
		if err != nil {
			return msg.UpdatedFail, false
		}
		return sb.StockLog()
	}
	if sh.StockType == global.Ex {
		return msg.ExGTStock, false
	}
	err = copier.Copy(&stock, &sb)
	if err != nil {
		return msg.Copier, false
	}
	stock.QTY = sb.InQTY + stock.QTY
	err = global.GDB.Create(&stock).Error
	if err != nil {
		return msg.UpdatedFail, false
	}
	return sb.StockLog()
}

func (sb *StockBody) StockLog() (m error, success bool) {
	if err := global.GDB.Create(&sb).Error; err != nil {
		return msg.UpdatedFail, false
	}
	return msg.UpdatedSuccess, true
}
