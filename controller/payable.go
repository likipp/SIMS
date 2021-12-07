package controller

import (
	"SIMS/internal/gins"
	"SIMS/models"
	"SIMS/services"
	"SIMS/utils"
	"SIMS/utils/msg"
	"github.com/gin-gonic/gin"
)

func CCreatePayBill(c *gin.Context) {
	var payable *models.Payable
	err := gins.ParseJSON(c, &payable)
	if err != nil {
		msg.Result(nil, err, 1, false, c)
		return
	}
	err, success := services.SCreatePayBill(payable)
	if success {
		msg.Result(nil, err, 0, true, c)
		return
	}
	msg.Result(nil, err, 1, false, c)
}

func CGetPayList(c *gin.Context) {
	param := utils.StringConvInt(c.Query("bill"))
	err, data, success := services.SGetPayList(param)
	if success {
		msg.Result(data, err, 0, true, c)
		return
	}
	msg.Result(nil, err, 1, false, c)
}

func CGetPayPie(c *gin.Context) {
	err, data, success := services.SGetPayPie()
	if success {
		msg.Result(data, err, 0, true, c)
		return
	}
	msg.Result(nil, err, 1, false, c)
}

func CGetExColumn(c *gin.Context) {
	err, data, success := services.SGetExColumn()
	if success {
		msg.Result(data, err, 0, true, c)
		return
	}
	msg.Result(nil, err, 1, false, c)
}

func CGetProductSale(c *gin.Context) {
	err, data, success := services.SGetProductSale()
	if success {
		msg.Result(data, err, 0, true, c)
		return
	}
	msg.Result(nil, err, 1, false, c)
}

func CGetProfit(c *gin.Context) {
	err, data, success := services.SGetProfit()
	if success {
		msg.Result(data, err, 0, true, c)
		return
	}
	msg.Result(nil, err, 1, false, c)
}

func CGetTotal(c *gin.Context) {
	err, data, success := services.SGetTotal()
	if success {
		msg.Result(data, err, 0, true, c)
		return
	}
	msg.Result(nil, err, 1, false, c)
}

func CGetCost(c *gin.Context) {
	err, data, success := services.SGetCost()
	if success {
		msg.Result(data, err, 0, true, c)
		return
	}
	msg.Result(nil, err, 1, false, c)
}
