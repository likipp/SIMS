package controller

import (
	"SIMS/internal/gins"
	"SIMS/models"
	"SIMS/services"
	"SIMS/utils/msg"
	"github.com/gin-gonic/gin"
)

func CCreateBrand(c *gin.Context) {
	var brand *models.Brand
	err := gins.ParseJSON(c, &brand)
	if err != nil {
		msg.Result(nil, err, 2, false, c)
		return
	}
	err, data, success := services.SCreateBrand(brand)
	if success {
		msg.Result(data, err, 0, true, c)
		return
	}
	msg.Result(nil, err, 2, false, c)
}

func CGetBrandSelectList(c *gin.Context) {
	var units []models.BrandSelect
	param := c.DefaultQuery("param", "")
	if param == "{}" {
		param = ""
	}
	err, units, success := services.SGetBrandSelectList(param)
	if !success {
		msg.Result(nil, err, 2, false, c)
		return
	}
	msg.Result(units, err, 0, true, c)
}

func CGetBrandTree(c *gin.Context) {
	err, list, success := services.SGetBrandTree()
	if !success {
		msg.Result(nil, err, 2, false, c)
		return
	}
	msg.Result(list, err, 0, true, c)
}
