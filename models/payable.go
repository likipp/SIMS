package models

import (
	"SIMS/global"
	"SIMS/internal/entity"
	"SIMS/utils/msg"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Payable struct {
	BaseModel
	Number     string  `json:"number" gorm:"comment:'单据号'"`
	SourceBill int     `json:"source_bill" gorm:"comment:'源单单号'"`
	ThisAmount float32 `json:"this_amount" gorm:"付款金额"`
	PayMethod  string  `json:"pay_method" gorm:"付款方式"`
}

type TotalAmount struct {
	Sum int `json:"sum"`
}

type PayList struct {
	SourceBill int     `json:"source_bill"`
	ThisAmount float32 `json:"this_amount"`
	PayMethod  string  `json:"pay_method"`
	CreatedAt *time.Time `json:"createdAt"`
	Status     int    `json:"status"`
}

func GetPayableDB() *gorm.DB {
	return entity.GetDBWithModel(global.GDB, new(Payable))
}

// CreatePayBill 创建应付单据
func (p *Payable) CreatePayBill() (error, bool) {
	var sh BillHeader
	var remainAmount float32
	// 生成应付单据单号
	p.Number = GenerateNumberWithYF()
	//err := global.GDB.Model(&BillEntry{}).Joins("JOIN bill_headers on bill_headers.id = bill_entries.header_id").Where( "bill_headers.id = ?", p.SourceBill).Pluck("bill_entries.total as total_amount", &p).Error
	// 查找需要关联的采购入库单据
	err := global.GDB.Where("bill_headers.id = ?", p.SourceBill).Find(&sh).Error
	if err != nil {
		return msg.GetFail, false
	}
	tx := global.GDB.Begin()
	// 计算应付金额
	remainAmount = sh.BillAmount - sh.RemainAmount - p.ThisAmount
	// 如果应付大于0, 说明还有欠款, 只需要更新入库单表头的已付金额
	if remainAmount > 0 {
		if err := tx.Model(&BillHeader{}).Where("id = ?", sh.ID).Update("remain_amount", sh.RemainAmount+p.ThisAmount).Error; err != nil {
			tx.Rollback()
			return msg.UpdatedFail, false
		}
		// 如果应付等于0, 说明货款将要结清, 需要更新入库单表头的已付金额, 并且关闭入库单状态
	} else if remainAmount == 0 {
		if err := tx.Model(&BillHeader{}).Where("id = ?", sh.ID).Updates(map[string]interface{}{"remain_amount": sh.BillAmount, "status": 1}).Error; err != nil {
			tx.Rollback()
			return msg.UpdatedFail, false
		}
		// 如果应付小于0, 说明货款已经结清, 不能继续付款
	} else {
		tx.Rollback()
		return msg.AmountSuccess, false
	}
	if err := tx.Create(&p).Error; err != nil {
		tx.Rollback()
		return msg.CreatedFail, false
	}
	tx.Commit()
	return msg.CreatedSuccess, true
}


func GetPayList(param int) (error, []PayList, bool) {
	db := GetPayableDB()
	var payList []PayList
	fmt.Println(param, "参数")
	if param != 0 {
		err := db.Select("payables.created_at,  payables.this_amount, bill_headers.status, payables.source_bill").Joins("JOIN bill_headers on payables.source_bill = bill_headers.id").Where("payables.source_bill = ?", param).Find(&payList).Error
		if err != nil {
			return msg.GetFail, payList, false
		}
		fmt.Println(payList, "错误消息")
		return msg.GetSuccess, payList, true
	}
	return msg.GetFail, payList, false
}