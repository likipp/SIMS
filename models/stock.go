package models

import (
	"SIMS/global"
	"SIMS/internal/entity"
	"SIMS/internal/schema"
	"SIMS/utils/msg"
	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
	"time"
)

type Stock struct {
	BaseModel
	PName     string `json:"p_name"`
	PNumber   string `json:"p_number"`
	WareHouse int    `json:"ware_house"`
	QTY       int    `json:"qty"`
}

type StockWithAction struct {
	Stock
	Action string `json:"action"`
}

type ExStock struct {
	Number     string     `json:"number"`
	CName      string     `json:"c_name"`
	CreatedAt  *time.Time `json:"created_at"`
	PayMethod  string     `json:"pay_method"`
	PNumber    string     `json:"p_number"`
	PName      string     `json:"p_name"`
	ExQTY      string     `json:"ex_qty"`
	UnitPrice  int        `json:"unit_price"`
	ExDiscount float32    `json:"ex_discount"`
	InDiscount float32    `json:"in_discount"`
	Cost       float32    `json:"cost"`
	Profit     float32    `json:"profit"`
	Total      float32    `json:"total"`
}

type InStock struct {
	ID       int                `json:"id"`
	Number    string            `json:"number"`
	Status   int                `json:"status"`
	CreatedAt *time.Time        `json:"created_at"`
	PayMethod string            `json:"pay_method"`
	PNumber   string            `json:"p_number"`
	PName     string            `json:"p_name"`
	InQTY     string            `json:"in_qty"`
	UnitPrice int               `json:"unit_price"`
	BillAmount   float32        `json:"bill_amount"`
	RemainAmount float32        `json:"remain_amount"`
}

type ExListQueryParams struct {
	schema.PaginationParam
	Sorter string `form:"sorter"`
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

func GetWareHouseQtyWithProduct(wareHouse int, product string) *Stock {
	var stock Stock
	var count int64
	db := GetStockDB()
	db.Where("ware_house = ? and p_number = ?", wareHouse, product).Find(&stock).Count(&count)
	if count > 0 {
		return &stock
	}
	return &stock
}

func (s *Stock) CheckStock(exQty int) (err error, success bool) {
	if s.QTY < exQty {
		return msg.ExGTStock, false
	}
	return nil, true
}

func (s *Stock) ChangeStock(action string, qty int) (error, bool) {
	db := GetStockDB()
	if err, success := s.CheckStock(qty); !success && action == global.Ex {
		return err, false
	}
	if action == global.Ex {
		s.QTY = s.QTY - qty
		if err := db.Model(s).Update("qty", s.QTY).Error; err != nil {
			return msg.UpdatedFail, false
		}
		return msg.UpdatedSuccess, true
	}
	s.QTY = s.QTY + qty
	if err := db.Model(s).Update("qty", s.QTY).Error; err != nil {
		return msg.UpdatedFail, false
	}
	return msg.UpdatedSuccess, true
}

func (s *Stock) CreateStockWithTransaction(tx *gorm.DB) (err error, success bool) {
	err = tx.Create(s).Error
	if err != nil {
		tx.Rollback()
		return msg.CreatedFail, false
	}
	return msg.CreatedSuccess, true
}

func (s *Stock) UpdateStockWithTransaction(tx *gorm.DB) (err error, success bool) {
	err = tx.Model(s).Update("qty", s.QTY).Error
	if err != nil {
		tx.Rollback()
		return msg.UpdatedFail, false
	}
	return msg.UpdatedSuccess, true
}

func GetStockList() (error, []Stock, bool) {
	var ss []Stock
	db := GetStockDB()
	if err := db.Find(&ss).Error; err != nil {
		return msg.GetFail, ss, false
	}
	return msg.GetSuccess, ss, true
}

func GetExStockList(params ExListQueryParams) (error, []ExStock, bool) {
	var el []ExStock
	db := global.GDB.Select("bill_headers.number, customs.c_name, bill_headers.created_at, bill_headers.pay_method, bill_headers.discount as h_discount, bill_entries.p_number, bill_entries.p_name, bill_entries.ex_qty, bill_entries.unit_price, bill_entries.discount as b_discount, bill_entries.total, bill_entries.cost, bill_entries.profit").Model(&BillHeader{}).Joins("left join bill_entries on bill_entries.header_id = bill_headers.id").Joins("left join customs on customs.id = bill_headers.custom").Where("bill_headers.stock_type = ?", "出库单")
	err := db.Scopes(schema.QueryPaging(params.PaginationParam)).Find(&el).Error
	//.Find(&el).Error
	if err != nil {
		return msg.GetFail, el, false
	}
	return msg.GetSuccess, el, true
}

func GetInStockList() (error, []InStock, bool) {
	var el []InStock
	err := global.GDB.Select("bill_headers.id, bill_headers.number, bill_headers.created_at, bill_headers.pay_method, bill_entries.p_number, bill_entries.p_name, bill_entries.in_qty, bill_entries.unit_price, bill_headers.bill_amount, bill_headers.remain_amount, bill_headers.status").Model(&BillHeader{}).Joins("left join bill_entries on bill_entries.header_id = bill_headers.id").Where("bill_headers.stock_type = ?", "入库单").Find(&el).Error
	if err != nil {
		return msg.GetFail, el, false
	}
	return msg.GetSuccess, el, true
}
