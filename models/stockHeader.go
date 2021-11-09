package models

import (
	"SIMS/global"
	"SIMS/utils/msg"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/copier"
)

// BillHeader 单据头
type BillHeader struct {
	BaseModel
	StockType string `json:"bill_type" gorm:"comment:'单据类型'"`
	Number    string `json:"bill_number" gorm:"comment:'单号'"`
	Custom    int    `json:"custom" gorm:"comment:'客户'"`
	//Supplier  int    `json:"supplier" gorm:"comment:'供应商'"`
	//Discount  int    `json:"discount"  gorm:"comment:'折扣'"`
	PayMethod    string  `json:"pay_method"  gorm:"comment:'收款方式'"`
	Status       int     `json:"status" gorm:"comment:'状态'"`
	BillAmount   float32 `json:"bill_amount" gorm:"订单金额"`
	RemainAmount float32 `json:"remain_amount" gorm:"剩余金额"`
}

// ExBillDetail 出库单详情
type ExBillDetail struct {
	Custom     string `json:"c_number"`
	CustomName string `json:"c_name"`
	BillHeader
	Body []BillEntry `json:"body"`
}

// InBillDetail 入库碟详情
type InBillDetail struct {
	BillHeader
	Body []BillEntry `json:"body"`
}

func (sh *BillHeader) Validate() error {
	err := validation.ValidateStruct(sh,
		validation.Field(&sh.Number, validation.Required.Error("单号不能为空")),
		validation.Field(&sh.StockType, validation.Required.Error("出入库类型不能为空")),
		validation.Field(&sh.PayMethod, validation.When(sh.StockType == global.Ex, validation.Required.Error("收款方式不能为空"))),
		validation.Field(&sh.Custom, validation.When(sh.StockType == global.Ex, validation.Required.Error("客户不能为空"))),
	)
	return err
}

func (sh *BillHeader) BillLog(sb []BillEntry) (err error, success bool) {
	var billTotal float32
	// 校验字段是否满足条件
	err = validation.Validate(sh, validation.NotNil)
	if err != nil {
		return err, false
	}
	// 开始数据库事务
	tx := global.GDB.Begin()
	// 创建单据表头信息
	if sh.StockType == "出库单" {
		sh.Status = 1
	} else {
		sh.Status = 0
		sh.RemainAmount = 0
	}
	if err = tx.Create(&sh).Error; err != nil {
		tx.Rollback()
		return msg.CreatedFail, false
	}
	// 循环表体明细, 根据单据类型, 更新库存数据
	for i := range sb {
		sb[i].HeaderID = sh.ID
		if sh.StockType == global.In {
			err, success = sb[i].InStockLog(tx)
			if !success {
				return err, false
			}
			billTotal += sb[i].Total
			continue
		}
		err, success = sb[i].ExStockLog(tx)
		if !success {
			return err, false
		}
	}
	err = tx.Model(&BillHeader{}).Where("number = ?", sh.Number).Update("bill_amount", billTotal).Error
	if err != nil {
		tx.Rollback()
		return msg.CreatedFail, false
	}
	tx.Commit()
	return err, true
}

// GetExBillDetail 获取出库订单详情信息
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

// GetInBillDetail 获取采购订单详情
func GetInBillDetail(number string) (error, InBillDetail, bool) {
	var header BillHeader
	var body []BillEntry
	var b InBillDetail
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
	b.Body = body
	return msg.GetSuccess, b, true
}
