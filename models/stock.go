package models

import (
	"SIMS/global"
	"SIMS/internal/entity"
	"SIMS/internal/schema"
	"SIMS/utils/msg"
	"fmt"
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
	ID         int        `json:"id" gorm:"index"`
	Number     string     `json:"number" gorm:"index"`
	CName      string     `json:"c_name" gorm:"index"`
	CreatedAt  *time.Time `json:"created_at" gorm:"index"`
	PayMethod  string     `json:"pay_method" gorm:"index"`
	PNumber    string     `json:"p_number" gorm:"index"`
	PName      string     `json:"p_name" gorm:"index"`
	ExQTY      string     `json:"ex_qty"`
	UnitPrice  int        `json:"unit_price"`
	ExDiscount float32    `json:"ex_discount"`
	InDiscount float32    `json:"in_discount"`
	Cost       float32    `json:"cost"`
	Profit     float32    `json:"profit"`
	Total      float32    `json:"total"`
}

type InStock struct {
	ID           int        `json:"id" gorm:"index"`
	Number       string     `json:"number" gorm:"index"`
	Status       int        `json:"status" gorm:"index"`
	CreatedAt    *time.Time `json:"created_at" gorm:"index"`
	PayMethod    string     `json:"pay_method" gorm:"index"`
	PNumber      string     `json:"p_number" gorm:"index"`
	PName        string     `json:"p_name" gorm:"index"`
	InQTY        string     `json:"in_qty"`
	UnitPrice    float32    `json:"unit_price"`
	BillAmount   float32    `json:"bill_amount"`
	RemainAmount float32    `json:"remain_amount"`
}

type InStockWithChild struct {
	ID           int        `json:"id"`
	Number       string     `json:"number"`
	Status       int        `json:"status"`
	CreatedAt    *time.Time `json:"created_at"`
	PayMethod    string     `json:"pay_method"`
	BillAmount   float32    `json:"bill_amount"`
	RemainAmount float32    `json:"remain_amount"`
	Child        []InChild  `json:"child"`
}

type InChild struct {
	PNumber   string  `json:"p_number" gorm:"index"`
	PName     string  `json:"p_name" gorm:"index"`
	InQTY     string  `json:"in_qty"`
	UnitPrice float32 `json:"unit_price"`
}

type InExStock struct {
	ID         int        `json:"id"`
	Number     string     `json:"number"`
	CName      string     `json:"c_name"`
	CreatedAt  *time.Time `json:"created_at"`
	PayMethod  string     `json:"pay_method"`
	PNumber    string     `json:"p_number"`
	PName      string     `json:"p_name"`
	ExQTY      int        `json:"ex_qty"`
	UnitPrice  int        `json:"unit_price"`
	ExDiscount float32    `json:"ex_discount"`
	InDiscount float32    `json:"in_discount"`
	Cost       float32    `json:"cost"`
	Profit     float32    `json:"profit"`
	Total      float32    `json:"total"`
	Status     int        `json:"status"`
	WareHouse  int        `json:"ware_house"`
	InQTY      int        `json:"in_qty"`
}

type InStockQueryParams struct {
	schema.PaginationParam
	Sorter    map[string]interface{} `form:"sorter"`
	Status    int                    `form:"status"`
	PNumber   string                 `form:"p_number"`
	PName     string                 `form:"p_name"`
	BeginTime *time.Time             `form:"begin_time"`
	EedTime   *time.Time             `form:"end_time"`
	Number    string                 `form:"number"`
	CreatedAt []string               `form:"created_at"`
}

type ExListQueryParams struct {
	schema.PaginationParam
	Sorter    map[string]interface{} `form:"sorter"`
	PNumber   string                 `form:"p_number"`
	PName     string                 `form:"p_name"`
	Number    string                 `form:"number"`
	PayMethod string                 `form:"pay_method"`
	//Custom    string     `form:"custom"`
	CustomName string   `form:"c_name"`
	CreatedAt  []string `form:"created_at"`
}

type StockQueryParams struct {
	Sorter    map[string]interface{} `form:"sorter"`
	PNumber   string                 `form:"p_number"`
	PName     string                 `form:"p_name"`
	WareHouse int                    `form:"ware_house"`
}

type InExStockQueryParams struct {
	schema.PaginationParam
	Sorter    map[string]interface{} `form:"sorter"`
	Status    int                    `form:"status"`
	PNumber   string                 `form:"p_number"`
	PName     string                 `form:"p_name"`
	BeginTime *time.Time             `form:"begin_time"`
	EedTime   *time.Time             `form:"end_time"`
	Number    string                 `form:"number"`
	WareHouse int                    `form:"ware_house"`
	CreatedAt []string               `form:"created_at"`
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

func GetWareHouseQtyWithProduct(wareHouse int, product string, tx *gorm.DB) *Stock {
	var stock Stock
	var count int64
	tx.Where("ware_house = ? and p_number = ?", wareHouse, product).Find(&stock).Count(&count)
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

func GetStockList(params StockQueryParams) (error, []Stock, bool) {
	var ss []Stock
	db := GetStockDB()
	if err := db.Where("qty != 0").Find(&ss).Error; err != nil {
		return msg.GetFail, ss, false
	}
	if v := params.PNumber; v != "" {
		if err := db.Where("p_number like ?", fmt.Sprintf("%s%s%s", "%", v, "%")).Error; err != nil {
			return msg.GetFail, nil, false
		}
	}
	if v := params.PName; v != "" {
		if err := db.Where("p_name like ?", fmt.Sprintf("%s%s%s", "%", v, "%")).Error; err != nil {
			return msg.GetFail, nil, false
		}
	}
	if v := params.WareHouse; v > 0 {
		if err := db.Where("ware_house = ?", params.WareHouse).Error; err != nil {
			return msg.GetFail, nil, false
		}
	}
	if err := db.Scopes(schema.QueryOrder(params.Sorter)).Find(&ss).Error; err != nil {
		return msg.GetFail, ss, false
	}
	return msg.GetSuccess, ss, true
}

func GetExStockList(params ExListQueryParams) (error, []ExStock, bool) {
	var el []ExStock
	//var queryScript = global.GDB.Select("header_id, sum(total) as total, sum(cost) as cost, sum(profit) as profit").Group("header_id")
	var selectColumns = "bill_headers.number, customs.c_name, bill_headers.created_at, bill_headers.pay_method, bill_entries.p_number, bill_entries.p_name, bill_entries.ex_qty, bill_entries.unit_price, cus.profit, cus.total, cus.cost"
	var jBEH = "left join bill_entries on bill_entries.header_id = bill_headers.id"
	var jCBH = "left join customs on customs.id = bill_headers.custom"
	var jQS = "join (select header_id, sum(total) as total, sum(cost) as cost, sum(profit) as profit from bill_entries group by header_id) as cus on cus.header_id = bill_entries.header_id"
	db := global.GDB.Select(selectColumns).Model(&BillHeader{}).Joins(jBEH).Joins(jCBH).Joins(jQS).Where("bill_headers.stock_type = ?", "出库单")
	err := db.Error
	if err != nil {
		return msg.GetFail, el, false
	}
	if v := params.PNumber; v != "" {
		if err = db.Where("p_number like ?", fmt.Sprintf("%s%s%s", "%", v, "%")).Error; err != nil {
			return msg.GetFail, nil, false
		}
	}
	if v := params.PName; v != "" {
		if err = db.Where("p_name like ?", fmt.Sprintf("%s%s%s", "%", v, "%")).Error; err != nil {
			return msg.GetFail, nil, false
		}
	}
	if v := params.Number; v != "" {
		if err = db.Where("number like ?", fmt.Sprintf("%s%s%s", "%", v, "%")).Error; err != nil {
			return msg.GetFail, nil, false
		}
	}
	if v := params.CustomName; v != "" {
		if err = db.Where("c_name like ?", fmt.Sprintf("%s%s%s", "%", v, "%")).Error; err != nil {
			return msg.GetFail, nil, false
		}
	}
	if v := params.PayMethod; v != "" {
		if err = db.Where("pay_method = ?", fmt.Sprintf("%s%s%s", "%", v, "%")).Error; err != nil {
			return msg.GetFail, nil, false
		}
	}
	if v := params.CreatedAt; len(v) > 0 {
		if err != nil {
			fmt.Println(err)
		}
		if err = db.Where("bill_headers.created_at BETWEEN ? AND ?", v[0], v[1]).Error; err != nil {
			return msg.GetFail, nil, false
		}
	}
	err = db.Scopes(schema.QueryOrder(params.Sorter)).Find(&el).Error
	if err != nil {
		return msg.GetFail, el, false
	}
	return msg.GetSuccess, el, true
}

func GetInStockList(params InStockQueryParams) (error, []InStockWithChild, bool) {
	var el []InStock
	var isc []InStockWithChild
	var hash = make(map[string]bool)
	var selectColumns = "bill_headers.id, bill_headers.number, bill_headers.created_at, bill_headers.pay_method, bill_entries.p_number, bill_entries.p_name, bill_entries.in_qty, bill_entries.unit_price, bill_headers.bill_amount, bill_headers.remain_amount, bill_headers.status"
	db := global.GDB.Select(selectColumns).Model(&BillHeader{}).Joins("left join bill_entries on bill_entries.header_id = bill_headers.id").Where("bill_headers.stock_type = ?", "入库单")
	err := db.Error
	if err != nil {
		return msg.GetFail, isc, false
	}
	if v := params.PNumber; v != "" {
		if err = db.Where("p_number like ?", fmt.Sprintf("%s%s%s", "%", v, "%")).Error; err != nil {
			return msg.GetFail, nil, false
		}
	}
	if v := params.PName; v != "" {
		if err = db.Where("p_name like ?", fmt.Sprintf("%s%s%s", "%", v, "%")).Error; err != nil {
			return msg.GetFail, nil, false
		}
	}
	if v := params.Status; v < 2 {
		if err = db.Where("status = ?", v).Error; err != nil {
			return msg.GetFail, nil, false
		}
	}
	if v := params.Number; v != "" {
		if err = db.Where("number like ?", fmt.Sprintf("%s%s%s", "%", v, "%")).Error; err != nil {
			return msg.GetFail, nil, false
		}
	}
	if v := params.CreatedAt; len(v) > 0 {
		if err != nil {
			fmt.Println(err)
		}
		if err = db.Where("bill_headers.created_at BETWEEN ? AND ?", v[0], v[1]).Error; err != nil {
			return msg.GetFail, nil, false
		}
	}
	if err = db.Scopes(schema.QueryOrder(params.Sorter)).Find(&el).Error; err != nil {
		return msg.GetFail, nil, false
	}

	for _, stock := range el {
		if !hash[stock.Number] {
			hash[stock.Number] = true
			i := InStockWithChild{
				ID:           stock.ID,
				Number:       stock.Number,
				Status:       stock.Status,
				CreatedAt:    stock.CreatedAt,
				PayMethod:    stock.PayMethod,
				BillAmount:   stock.BillAmount,
				RemainAmount: stock.RemainAmount,
				Child:        nil,
			}
			isc = append(isc, i)
		}
	}
	for i := 0; i < len(isc); i++ {
		for y := 0; y < len(el); y++ {
			if isc[i].Number == el[y].Number {
				isc[i].Child = append(isc[i].Child, InChild{
					PNumber:   el[y].PNumber,
					PName:     el[y].PName,
					InQTY:     el[y].InQTY,
					UnitPrice: el[y].UnitPrice,
				})
			}
		}
	}
	fmt.Println(isc, "入库列表")
	return msg.GetSuccess, isc, true
}

func GetInExStockList(params InExStockQueryParams) (error, []InExStock, bool) {
	var ies []InExStock
	var selectColumns = "bill_headers.id, bill_headers.number, bill_headers.created_at, bill_entries.p_number, bill_entries.p_name, bill_entries.ware_house, bill_entries.in_qty, bill_entries.ex_qty, bill_entries.ex_discount, bill_entries.in_discount, bill_entries.unit_price, bill_entries.total"
	db := global.GDB.Select(selectColumns).Model(&BillHeader{}).Joins("left join bill_entries on bill_entries.header_id = bill_headers.id").Order("bill_headers.created_at")
	err := db.Error
	if err != nil {
		return msg.GetFail, ies, false
	}
	if v := params.PNumber; v != "" {
		if err = db.Where("p_number like ?", fmt.Sprintf("%s%s%s", "%", v, "%")).Error; err != nil {
			return msg.GetFail, nil, false
		}
	}
	if v := params.PName; v != "" {
		if err = db.Where("p_name like ?", fmt.Sprintf("%s%s%s", "%", v, "%")).Error; err != nil {
			return msg.GetFail, nil, false
		}
	}
	if v := params.WareHouse; v > 0 {
		if err = db.Where("ware_house = ?", v).Error; err != nil {
			return msg.GetFail, nil, false
		}
	}
	if err = db.Scopes(schema.QueryOrder(params.Sorter)).Find(&ies).Error; err != nil {
		return msg.GetFail, nil, false
	}
	return msg.GetSuccess, ies, true
}
