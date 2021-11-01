package controller

import (
	"SIMS/internal/gins"
	"SIMS/models"
	"SIMS/services"
	"SIMS/utils/msg"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func CCreateWareHouse(c *gin.Context) {
	var wareHouse *models.WareHouse
	err := gins.ParseJSON(c, &wareHouse)
	if err != nil {
		msg.Result(nil, err, 2, false, c)
		return
	}
	//err = validation.Validate(wareHouse, validation.NotNil)
	err = validation.Validate(wareHouse, validation.NotNil)
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

func CGetWareHouseList(c *gin.Context) {
	var products []models.WareHouseSelect
	param := c.DefaultQuery("param", "")
	if param == "{}" {
		param = ""
	}
	err, products, success := services.SGetWareHouseList(param)
	if !success {
		msg.Result(nil, err, 2, false, c)
		return
	}
	msg.Result(products, err, 1, true, c)
}

func CWareHouseList(c *gin.Context) {
	var products []models.WareHouseSelect
	param := c.DefaultQuery("param", "")
	if param == "{}" {
		param = ""
	}
	err, products, success := services.SGetWareHouseList(param)
	if !success {
		msg.Result(nil, err, 2, false, c)
		return
	}
	msg.Result(products, err, 1, true, c)
}
