package controller

import (
	"SIMS/internal/gins"
	"SIMS/models"
	"SIMS/services"
	"SIMS/utils/msg"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// Bill 定义前端传递过来的单据信息
type Bill struct {
	models.BaseModel
	Status    int         `json:"status"`
	StockType string      `json:"bill_type"`
	Number    string      `json:"bill_number"`
	Custom    int         `json:"custom"`
	Discount  int         `json:"discount"`
	PayMethod string      `json:"pay_method"`
	Body      []StockBody `json:"body"`
}

type StockBody struct {
	ID         int     `json:"id"`
	HeaderID   int     `json:"header_id"`
	PNumber    string  `json:"p_number"`
	PName      string  `json:"p_name"`
	WareHouse  int     `json:"ware_house"`
	ExQTY      int     `json:"ex_qty"`
	InQTY      int     `json:"in_qty"`
	UnitPrice  float32 `json:"unit_price"`
	Discount   float32 `json:"discount"`
	ExDiscount float32 `json:"ex_discount"`
	InDiscount float32 `json:"in_discount"`
	Cost       float32 `json:"cost"`
	Profit     float32 `json:"profit"`
	Total      float32 `json:"total"`
}

func CStockHeader(c *gin.Context) {
	var sb []models.BillEntry
	var sh models.BillHeader
	var stock Bill
	err := gins.ParseJSON(c, &stock)
	if err != nil {
		msg.Result(nil, msg.QueryParamsFail, 1, false, c)
		return
	}
	// 复制前端传递胡单据头信息到 sh BillHeader
	if err = copier.Copy(&sh, stock); err != nil {
		msg.Result(nil, msg.Copier, 1, false, c)
		return
	}
	// 复制前端传递胡单据体信息到 sb BillEntry
	if err = copier.Copy(&sb, stock.Body); err != nil {
		msg.Result(nil, msg.Copier, 1, false, c)
		return
	}
	err, success := services.SStockHeader(sh, sb)
	if success {
		msg.Result(nil, err, 0, true, c)
		return
	}
	msg.Result(nil, err, 1, false, c)
	return
}

func CDeleteBill(c *gin.Context) {
	number := c.Param("number")
	fmt.Println(number, "number")
	err, success := services.SDeleteBillByID(number)
	if success {
		msg.Result(nil, err, 0, true, c)
		return
	}
	msg.Result(nil, err, 1, false, c)
}

func CUpdateBillByID(c *gin.Context) {
	var stock Bill
	var sb []models.BillEntry
	var sh models.BillHeader
	err := gins.ParseJSON(c, &stock)
	fmt.Println(stock, "body")
	if err != nil {
		msg.Result(nil, msg.QueryParamsFail, 1, false, c)
		return
	}
	if err = copier.Copy(&sh, stock); err != nil {
		msg.Result(nil, msg.Copier, 1, false, c)
		return
	}
	// 复制前端传递到单据体信息到 sb BillEntry
	if err = copier.Copy(&sb, stock.Body); err != nil {
		msg.Result(nil, msg.Copier, 1, false, c)
		return
	}
	err, success := services.SUpdateBillByID(sh.Number, sb)
	fmt.Println(err, "error错误消息", success)
	if success {
		msg.Result(nil, err, 0, true, c)
		return
	}
	msg.Result(nil, err, 1, false, c)
}

func CGetExBillDetail(c *gin.Context) {
	number := c.Query("number")
	err, data, success := services.SGetExBillDetail(number)
	if success {
		msg.Result(data, msg.GetSuccess, 1, true, c)
		return
	}
	msg.Result(nil, err, 1, false, c)
	return
}

func CGetInBillDetail(c *gin.Context) {
	number := c.Query("number")
	err, data, success := services.SGetInBillDetail(number)
	if success {
		msg.Result(data, msg.GetSuccess, 0, true, c)
		return
	}
	msg.Result(nil, err, 1, false, c)
	return
}
