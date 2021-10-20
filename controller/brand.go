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
		msg.Result(data, err, 1, true, c)
		return
	}
	msg.Result(nil, err, 2, false, c)
}
