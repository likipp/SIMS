package models

import (
	"SIMS/global"
	"SIMS/utils/msg"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/copier"
)

type BillHeader struct {
	BaseModel
	StockType string `json:"bill_type" gorm:"comment:'单据类型'"`
	Number    string `json:"bill_number" gorm:"comment:'单号'"`
	Custom    int    `json:"custom" gorm:"comment:'客户'"`
	//Supplier  int    `json:"supplier" gorm:"comment:'供应商'"`
	Discount  int    `json:"discount"  gorm:"comment:'折扣'"`
	PayMethod string `json:"pay_method"  gorm:"comment:'收款方式'"`
}

type ExBillDetail struct {
	Custom     string `json:"c_number"`
	CustomName string `json:"c_name"`
	BillHeader
	Body []BillEntry `json:"body"`
}

func (sh *BillHeader) Validate() error {
	err := validation.ValidateStruct(sh,
		validation.Field(&sh.Number, validation.Required.Error("单号不能为空")),
		validation.Field(&sh.StockType, validation.Required.Error("出入库类型不能为空")),
		validation.Field(&sh.PayMethod, validation.Required.Error("收款方式不能为空")),
		validation.Field(&sh.Custom, validation.When(sh.StockType == global.Ex, validation.Required.Error("客户不能为空"))),
		//validation.Field(&sh.Supplier, validation.When(sh.StockType == global.In, validation.Required.Error("供应商不能为空"))),
	)
	return err
}

func (sh *BillHeader) BillLog(sb []BillEntry) (err error, success bool) {
	err = validation.Validate(sh, validation.NotNil)
	if err != nil {
		return err, false
	}
	tx := global.GDB.Begin()
	if err = tx.Create(&sh).Error; err != nil {
		tx.Rollback()
		return msg.CreatedFail, false
	}
	for i := range sb {
		sb[i].HeaderID = sh.ID
		if sh.StockType == global.In {
			err, success = sb[i].InStockLog(tx)
			if !success {
				return err, false
			}
			continue
		}
		err, success = sb[i].ExStockLog(tx)
		if !success {
			return err, false
		}
	}
	tx.Commit()
	return err, true
}

// GetExBillDetail 获取订单详情信息
func GetExBillDetail(number string) (error, ExBillDetail, bool) {
	var header BillHeader
	var body []BillEntry
	var b ExBillDetail
	var c Custom
	err := global.GDB.Where("number = ?", number).Find(&header).Error
	if err != nil {
		return msg.GetFail, b, false
	}
	global.GDB.Where("id = ?", header.Custom).Find(&c)
	if err != nil {
		return msg.GetFail, b, false
	}
	err = global.GDB.Where("header_id = ?", header.ID).Find(&body).Error
	if err != nil {
		return msg.GetFail, b, false
	}
	if err = copier.Copy(&b, &header); err != nil {
		return msg.Copier, b, false
	}
	b.Custom = c.CNumber
	b.CustomName = c.CName
	b.Body = body
	return msg.GetSuccess, b, true
}
