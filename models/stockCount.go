package models

import (
	"SIMS/global"
	"SIMS/utils"
	"SIMS/utils/msg"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/copier"
	"time"
)

type StockCount struct {
	BaseModel
	PNumber   string `json:"p_number" gorm:"comment:'产品代码'"`
	PName     string `json:"p_name" gorm:"comment:'产品名称'"`
	ExQTY     int    `json:"ex_qty" gorm:"comment:'出库数量'"`
	InQTY     int    `json:"in_qty" gorm:"comment:'入库数量'"`
	WareHouse int    `json:"ware_house" gorm:"comment:'仓库'"`
	StockType string `json:"type" gorm:"comment:'单据类型'"`
	Number    string `json:"number" gorm:"comment:'单号'"`
	Custom    int    `json:"custom" gorm:"comment:'客户'"`
	Courier   int    `json:"courier" gorm:"comment:'供应商'"`
}

func (s *StockCount) Validate() error {
	err := validation.ValidateStruct(s,
		validation.Field(&s.PNumber, validation.Required.Error("产品代码不能为空")),
		validation.Field(&s.PName, validation.Required.Error("产品名称不能为空")),
		validation.Field(&s.WareHouse, validation.Required.Error("仓库不能为空")),
		validation.Field(&s.StockType, validation.Required.Error("出入库类型为空")),
		validation.Field(&s.InQTY, validation.When(s.StockType == global.In, validation.Required.Error("数量不能为空"))),
		validation.Field(&s.ExQTY, validation.When(s.StockType != global.In, validation.Required.Error("数量不能为空"))),
		validation.Field(&s.Custom, validation.When(s.StockType == global.In, validation.Required.Error("供应商不能为空"))),
		validation.Field(&s.Courier, validation.When(s.StockType != global.In, validation.Required.Error("客户不能为空"))),
	)
	return err
}

func (s *StockCount) StockCount() (err error, success bool) {
	stock := GetWareHouseQtyWithProduct(s.WareHouse, s.PNumber)
	if stock.QTY > 0 {
		if s.StockType == global.Ex {
			if stock.QTY < s.ExQTY {
				return msg.ExGTStock, false
			}
			stock.QTY = stock.QTY - s.ExQTY
			err = global.GDB.Model(stock).Update("qty", stock.QTY).Error
			if err != nil {
				return msg.UpdatedFail, false
			}
			return s.StockLog()
		}
		stock.QTY = s.InQTY + stock.QTY
		err = global.GDB.Model(stock).Update("qty", stock.QTY).Error
		if err != nil {
			return msg.UpdatedFail, false
		}
		return s.StockLog()
		//return msg.UpdatedSuccess, true
	}
	if s.StockType == global.Ex {
		return msg.ExGTStock, false
	}
	copier.Copy(&stock, &s)
	stock.QTY = s.InQTY + stock.QTY
	err = global.GDB.Create(&stock).Error
	if err != nil {
		return msg.UpdatedFail, false
	}
	return s.StockLog()
	//return msg.UpdatedSuccess, true
}

func (s *StockCount) StockLog() (m error, success bool) {
	//fmt.Println(s, "库存")
	n := Number(s.StockType)
	s.Number = n
	if err := global.GDB.Create(&s).Error; err != nil {
		return msg.UpdatedFail, false
	}
	return msg.UpdatedSuccess, true
}

func Number(t string) (n string) {
	var s StockCount
	var total int64
	time := time.Now().Format("20060102")
	timeParam := fmt.Sprintf("%s%s%s", "%", time, "%")
	global.GDB.Raw("select number from stock_counts where stock_type = ? and number like ? order by number desc limit 1", t, timeParam).Scan(&s).Count(&total)
	if total > 0 {
		number := s.Number
		lastString := number[len(number)-3:]
		newNumStr := utils.IntConvJoin(lastString)
		if t == global.Ex {
			return fmt.Sprintf("%s%s%s", "EX", time, newNumStr)
		}
		if t == global.In {
			return fmt.Sprintf("%s%s%s", "In", time, newNumStr)
		}
	}
	if t == global.Ex {
		return fmt.Sprintf("%s%s%s", "EX", time, "01")
	}
	//if t == global.In {
	//
	//}
	return fmt.Sprintf("%s%s%s", "In", time, "01")
}
