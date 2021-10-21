package models

import (
	"SIMS/global"
	"SIMS/utils/msg"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/copier"
)

type StockCount struct {
	BaseModel
	PNumber         string       `json:"p_number"`
	PName           string       `json:"p_name"`
	Qty             int          `json:"qty"`
	WareHouse       int          `json:"ware_house"`
	StockType       string       `json:"type"` //单据类型
	Number          string       `json:"number"`
}

func (s *StockCount) Validate() error {
	err := validation.ValidateStruct(s,
		validation.Field(&s.PNumber, validation.Required.Error("产品代码不能为空")),
		validation.Field(&s.PName, validation.Required.Error("产品名称不能为空")),
		validation.Field(&s.WareHouse, validation.Required.Error("仓库不能为空")),
		validation.Field(&s.Qty, validation.Required.Error("数量不能为空")),
		validation.Field(&s.StockType, validation.Required.Error("出入库类型为空")),
		)
	return err
}

func (s *StockCount) StockCount() (err error, success bool) {
	stock, ok := GetWareHouseQtyWithProduct(s.WareHouse, s.PNumber)
	fmt.Println(stock, ok)
	if ok {
		if s.StockType == global.Ex {
			if stock.QTY < s.Qty {
				return msg.ExGTStock, false
			}
			s.Qty = s.Qty - stock.QTY
			err = global.GDB.Model(stock).Update("qty", s.Qty).Error
			if err != nil {
				return msg.UpdatedFail, false
			}
			return msg.UpdatedSuccess, true
		}
		s.Qty = s.Qty + stock.QTY
		err = global.GDB.Model(stock).Update("qty", s.Qty).Error
		if err != nil {
			return msg.UpdatedFail, false
		}
		return msg.UpdatedSuccess, true
	}
	if s.StockType == global.Ex {
		//if stock.QTY < s.Qty {
		//	return msg.ExGTStock, false
		//}
		return msg.ExGTStock, false
		//s.Qty = s.Qty - stock.QTY
		//err = global.GDB.Create(&s).Error
		//if err != nil {
		//	return msg.UpdatedFail, false
		//}
		//return msg.UpdatedSuccess, true
	}
	//s.Qty = s.Qty + stock.QTY
	copier.Copy(&stock, &s)
	fmt.Println(stock, "stock", s)
	err = global.GDB.Create(&s).Error
	if err != nil {
		return msg.UpdatedFail, false
	}
	return msg.UpdatedSuccess, true
}