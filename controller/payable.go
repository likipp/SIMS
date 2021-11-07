package controller

import (
	"SIMS/internal/gins"
	"SIMS/models"
	"SIMS/services"
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
