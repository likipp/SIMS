package controller

import (
	"SIMS/global"
	"SIMS/internal/gins"
	"SIMS/models"
	"SIMS/services"
	"SIMS/utils"
	"SIMS/utils/msg"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type StockChange struct {
	Action string `json:"action"`
	QTY    int    `json:"qty"`
}

//func CInStock(c *gin.Context) {
//	var stock *models.Stock
//	err := gins.ParseJSONWithPath(c, &stock)
//	if err != nil {
//		msg.Result(nil, err, 2, false, c)
//		return
//	}
//	err = validation.Validate(stock, validation.NotNil)
//	if err != nil {
//		msg.Result(nil, err, 2, false, c)
//		return
//	}
//	err, data, success := services.SInStock(stock)
//	if success {
//		msg.Result(data, err, 1, true, c)
//		return
//	}
//	msg.Result(nil, err, 2, false, c)
//}
//
//func CExStock(c *gin.Context) {
//	var stock *models.Stock
//	err := gins.ParseJSON(c, &stock)
//	if err != nil {
//		msg.Result(nil, err, 2, false, c)
//		return
//	}
//	err = validation.Validate(stock, validation.NotNil)
//	if err != nil {
//		msg.Result(nil, err, 2, false, c)
//		return
//	}
//	err, data, success := services.SExStock(stock)
//	if success {
//		msg.Result(data, err, 1, true, c)
//		return
//	}
//	msg.Result(nil, err, 2, false, c)
//}

func CInOrExStock(c *gin.Context) {
	var stock *models.Stock
	err, action := gins.ParamQuery(c, "aa")
	err = gins.ParseJSON(c, &stock)
	if err != nil {
		msg.Result(nil, err, 2, false, c)
		return
	}
	err = validation.Validate(stock, validation.NotNil)
	if err != nil {
		msg.Result(nil, err, 2, false, c)
		return
	}
	if action == global.In {
		msg.Result(nil, err, 2, false, c)
		return
	}
	msg.Result(nil, err, 2, false, c)
}

func CGetStockList(c *gin.Context) {
	var params models.StockQueryParams
	err := gins.ParseQuery(c, &params)
	if err != nil {
		msg.Result(nil, err, 1, false, c)
		return
	}
	err, i, success := services.SGetStockList(params)
	if !success {
		msg.Result(i, err, 2, false, c)
		return
	}
	msg.Result(i, err, 1, true, c)
}

func CSChangeStock(c *gin.Context) {
	var change StockChange
	var s models.Stock
	id := utils.StringConvInt(c.Param("id"))
	if err := global.GDB.Where("id = ?", id).Find(&s).Error; err != nil {
		msg.Result(nil, msg.GetFail, 2, false, c)
		return
	}
	if err := gins.ParseJSON(c, &change); err != nil {
		msg.Result(nil, err, 2, false, c)
		return
	}
	err, success := services.SChangeStock(s, change.Action, change.QTY)
	if !success {
		msg.Result(nil, err, 2, false, c)
		return
	}
	msg.Result(nil, err, 1, true, c)
}

func CGetExStockList(c *gin.Context) {
	var params models.ExListQueryParams
	err := gins.ParseQuery(c, &params)
	if err != nil {
		msg.Result(nil, err, 1, false, c)
		return
	}

	err, list, success := services.SGetExStockList(params)
	if !success {
		msg.Result(nil, err, 1, false, c)
		return
	}
	msg.ResultWithPageInfo(list, err, 0, success, int64(len(list)), params.Current, params.PageSize, c)
}

func CGetInStockList(c *gin.Context) {
	var params models.InStockQueryParams
	err := gins.ParseQuery(c, &params)
	if err != nil {
		msg.Result(nil, err, 1, false, c)
		return
	}
	err, list, success := services.SGetInStockList(params)
	if !success {
		msg.Result(nil, err, 2, false, c)
		return
	}
	msg.Result(list, err, 1, true, c)
}
