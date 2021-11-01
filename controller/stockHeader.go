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

type ExBill struct {
	models.BaseModel
	StockType     string       `json:"bill_type"`
	Number    string       `json:"bill_number"`
	Custom        int          `json:"custom"`
	Discount      int          `json:"discount"`
	PayMethod     string       `json:"pay_method"`
	Body          []StockBody  `json:"body"`
}

type StockBody struct {
	PNumber   string  `json:"p_number"`
	PName     string  `json:"p_name"`
	WareHouse int     `json:"ware_house"`
	ExQTY     int     `json:"ex_qty"`
	InQTY     int     `json:"in_qty"`
	UnitPrice float32 `json:"unit_price"`
	Discount  float32 `json:"discount"`
	Total     float32 `json:"total"`
}

func CStockHeader(c *gin.Context) {
	var sb []models.BillEntry
	var sh models.BillHeader
	var stock ExBill
	err := gins.ParseJSON(c, &stock)
	if err != nil {
		msg.Result(nil, msg.QueryParamsFail, 2, false, c)
		return
	}
	fmt.Println(stock, "stock")
	if err = copier.Copy(&sh, stock); err != nil {
		msg.Result(nil, msg.Copier, 2, false, c)
		return
	}
	fmt.Println(sh, "单据头")
	if err = copier.Copy(&sb, stock.Body); err != nil {
		msg.Result(nil, msg.Copier, 2, false, c)
		return
	}
	err, success := services.SStockHeader(sh, sb)
	if success {
		msg.Result(nil, msg.UpdatedSuccess, 1, true, c)
		return
	}
	msg.Result(nil, err, 1, false, c)
	return
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
