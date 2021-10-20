package controller

import (
	"SIMS/internal/gins"
	"SIMS/models"
	"SIMS/services"
	"SIMS/utils/msg"
	"github.com/gin-gonic/gin"
)

func CCreateWareHouse(c *gin.Context) {
	var wareHouse *models.WareHouse
	err := gins.ParseJSON(c, &wareHouse)
	if err != nil {
		msg.Result(nil, err, 2, false, c)
		return
	}
	err, data, success := services.SCreateWareHouse(wareHouse)
	if success {
		msg.Result(data, err, 1, true, c)
		return
	}
	msg.Result(nil, err, 2, false, c)
}