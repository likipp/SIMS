package controller

import (
	"SIMS/models"
	"SIMS/services"
	"SIMS/utils"
	"SIMS/utils/msg"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ACreateCourier(c *gin.Context) {
	var courier *models.Courier
	err := c.BindJSON(&courier)
	err, ok := err.(validator.ValidationErrors)
	if !ok {
		msg.Result(nil, err, 0, false, c)
		return
	}
	//_ = c.ShouldBindBodyWith(&courier, binding.JSON).Error()
	err, data := services.SCreateCourier(courier)
	if err != nil {
		msg.Result(nil, err, 0, false, c)
		return
	}
	msg.Result(data, err, 0, true, c)
}

func AUpdateCourier(c *gin.Context) {
	var r *models.Courier
	_ = c.ShouldBindJSON(&r)
	id := c.Param("id")
	err := services.SUpdateCourier(r, utils.StringConvInt(id))
	if err != nil {
		msg.Result(nil, err, 0, false, c)
	} else {
		msg.Result(nil, err, 0, true, c)
	}
}

func ADeleteCourier(c *gin.Context) {
	var r models.Courier
	id := c.Param("id")
	r.ID = utils.StringConvInt(id)
	err := services.SDeleteCourier(&r)
	if err != nil {
		msg.Result(nil, err, 0, false, c)
	} else {
		msg.Result(nil, err, 0, true, c)
	}
}
