package controller

import (
	"SIMS/internal/gins"
	"SIMS/models"
	"SIMS/services"
	"SIMS/utils/msg"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func CCreateProduct(c *gin.Context) {
	var product *models.Products
	err := gins.ParseJSON(c, &product)
	if err != nil {
		msg.Result(nil, err, 2, false, c)
		return
	}
	err = validation.Validate(product, validation.NotNil)
	if err != nil {
		msg.Result(nil, err, 2, false, c)
		return
	}
	err, data, success := services.SCreateProduct(product)
	if success {
		msg.Result(data, err, 1, true, c)
		return
	}
	msg.Result(nil, err, 2, false, c)
}
