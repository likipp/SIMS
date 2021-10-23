package controller

import (
	"SIMS/internal/gins"
	"SIMS/models"
	"SIMS/services"
	"SIMS/utils/msg"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type Stock struct {
	models.BaseModel
	StockType string `json:"type"`
	Number    string `json:"number"`
	Custom    int    `json:"custom"`
	//Courier         int           `json:"courier"`
	Discount  int         `json:"discount"`
	PayMethod string      `json:"pay_method"`
	Body      []StockBody `json:"body"`
}

type StockBody struct {
	//HeaderID        int          `json:"header_id"`
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
	var sb []models.StockBody
	var sh models.StockHeader
	var stock Stock
	err := gins.ParseJSON(c, &stock)
	if err != nil {
		msg.Result(nil, msg.QueryParamsFail, 2, false, c)
		return
	}
	if err = copier.Copy(&sh, stock); err != nil {
		msg.Result(nil, msg.Copier, 2, false, c)
		return
	}
	if err = copier.Copy(&sb, stock.Body); err != nil {
		msg.Result(nil, msg.Copier, 2, false, c)
		return
	}
	err, success := services.SStockHeader(sh, sb)
	if success {
		msg.Result(nil, msg.UpdatedSuccess, 1, true, c)
		return
	}
	msg.Result(nil, msg.UpdatedFail, 1, true, c)
	return
}
