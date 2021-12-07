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
	SourceBill int        `json:"source_bill"`
	ThisAmount float32    `json:"this_amount"`
	PayMethod  string     `json:"pay_method"`
	CreatedAt  *time.Time `json:"createdAt"`
	Status     int        `json:"status"`
}

type PayPie struct {
	Type  string  `json:"type"`
	Value float32 `json:"value"`
}

type ExColumn struct {
	Month int     `json:"month"`
	Brand string  `json:"brand"`
	Value float32 `json:"value"`
}

type ProfitCompare struct {
	ThisMonth float32 `json:"this_month"`
	PreMonth  float32 `json:"pre_month"`
	Up        bool    `json:"up"`
}

type ProductSale struct {
	Product string  `json:"product"`
	Value   float32 `json:"value"`
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
	if param != 0 {
		err := db.Select("payables.pay_method, payables.created_at,  payables.this_amount, bill_headers.status, payables.source_bill").Joins("JOIN bill_headers on payables.source_bill = bill_headers.id").Where("payables.source_bill = ?", param).Find(&payList).Error
		if err != nil {
			return msg.GetFail, payList, false
		}
		fmt.Println(payList, "错误消息")
		return msg.GetSuccess, payList, true
	}
	return msg.GetFail, payList, false
}

func GetPayPie() (error, []PayPie, bool) {
	var payPie []PayPie
	var s = "brands.name as type, bill_headers.bill_amount - bill_headers.remain_amount as value"
	var j1 = "join bill_entries on bill_entries.header_id = bill_headers.id"
	var j2 = "join products on products.p_number = bill_entries.p_number"
	var j3 = "join brands on brands.id = products.brand"
	if err := global.GDB.Table("bill_headers").Select(s).Joins(j1).Joins(j2).Joins(j3).Where("bill_headers.stock_type = ?", "入库单").Group("brands.name").Scan(&payPie).Error; err != nil {
		return msg.GetFail, payPie, false
	}
	return msg.GetSuccess, payPie, true
}

func GetExColumn() (error, []ExColumn, bool) {
	var ec []ExColumn
	var s = "DATE_FORMAT(bill_headers.created_at,'%m') as month, brands.name as brand, sum(bill_entries.total) as value"
	var j1 = "join bill_entries on bill_headers.id = bill_entries.header_id"
	var j2 = "join products on bill_entries.p_number = products.p_number"
	var j3 = "join brands on products.brand = brands.id"
	if err := global.GDB.Table("bill_headers").Select(s).Joins(j1).Joins(j2).Joins(j3).Where("bill_headers.stock_type = ?", "出库单").Group("DATE_FORMAT(bill_headers.created_at,'%m')").Group("products.brand").Order("DATE_FORMAT(bill_headers.created_at,'%m')").Scan(&ec).Error; err != nil {
		return msg.GetFail, ec, false
	}
	return msg.GetSuccess, ec, true
}

func GetProductSale() (error, []ProductSale, bool) {
	var ps []ProductSale
	var s = "bill_entries.p_name as product, sum(bill_entries.total) as value"
	var j1 = "join bill_entries on bill_headers.id = bill_entries.header_id"
	if err := global.GDB.Table("bill_headers").Select(s).Joins(j1).Where("bill_headers.stock_type = ?", "出库单").Group("bill_entries.p_number").Order("sum(bill_entries.total) desc").Limit(5).Scan(&ps).Error; err != nil {
		return msg.GetFail, ps, false
	}
	return msg.GetSuccess, ps, true
}

func GetProfit() (error, ProfitCompare, bool) {
	var pc ProfitCompare
	pc.Up = false
	month := time.Now().Month()
	var j1 = "join bill_entries on bill_headers.id = bill_entries.header_id"
	if err := global.GDB.Table("bill_headers").Select("sum(bill_entries.profit) as this_month").Joins(j1).Where("bill_headers.stock_type = ? and DATE_FORMAT(bill_headers.created_at,'%m') = ?", "出库单", month).Scan(&pc).Error; err != nil {
		return msg.GetFail, pc, false
	}
	if err := global.GDB.Table("bill_headers").Select("sum(bill_entries.profit) as pre_month").Joins(j1).Where("bill_headers.stock_type = ? and DATE_FORMAT(bill_headers.created_at,'%m') = ?", "出库单", month-1).Scan(&pc).Error; err != nil {
		return msg.GetFail, pc, false
	}
	if pc.ThisMonth > pc.PreMonth {
		pc.Up = true
	}
	return msg.GetSuccess, pc, true
}

func GetTotal() (error, float32, bool) {
	var total float32
	var j1 = "join bill_entries on bill_headers.id = bill_entries.header_id"
	if err := global.GDB.Table("bill_headers").Select("sum(bill_entries.total)").Joins(j1).Where("bill_headers.stock_type = ?", "出库单").Scan(&total).Error; err != nil {
		return msg.GetFail, 0, false
	}
	return msg.GetSuccess, total, true
}

func GetCost() (error, float32, bool) {
	var total float32
	var j1 = "join bill_entries on bill_headers.id = bill_entries.header_id"
	if err := global.GDB.Table("bill_headers").Select("sum(bill_entries.total)").Joins(j1).Where("bill_headers.stock_type = ?", "入库单").Scan(&total).Error; err != nil {
		return msg.GetFail, 0, false
	}
	return msg.GetSuccess, total, true
}
