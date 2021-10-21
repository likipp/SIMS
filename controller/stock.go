package controller

import (
	"SIMS/internal/gins"
	"SIMS/models"
	"SIMS/utils/msg"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

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
	err, action := gins.ParamQuery(c)
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
	if action == "in" {
		//err, data, success := services.SInStock(stock)
		//if success {
		//	msg.Result(data, err, 1, true, c)
		//	return
		//}
		msg.Result(nil, err, 2, false, c)
		return
	}
	//err, data, success := services.SExStock(stock)
	//if success {
	//	msg.Result(data, err, 1, true, c)
	//	return
	//}
	msg.Result(nil, err, 2, false, c)
}