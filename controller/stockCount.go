package controller

import (
	"SIMS/internal/gins"
	"SIMS/models"
	"SIMS/services"
	"SIMS/utils/msg"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func CStockCount(c *gin.Context) {
	var stockCount *models.StockCount
	err := gins.ParseJSON(c, &stockCount)
	if err != nil {
		msg.Result(nil, err, 2, false, c)
		return
	}
	err = validation.Validate(stockCount, validation.NotNil)
	if err != nil {
		msg.Result(nil, err, 2, false, c)
		return
	}
	err, success := services.SStockCount(stockCount)
	if success {
		msg.Result(nil, err, 1, true, c)
		return
	}
	msg.Result(nil, err, 2, false, c)
}