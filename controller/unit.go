package controller

import (
	"SIMS/internal/gins"
	"SIMS/models"
	"SIMS/services"
	"SIMS/utils/msg"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

func CCreateUnit(c *gin.Context) {
	var unit *models.Unit
	err := gins.ParseJSON(c, &unit)
	if err != nil {
		msg.Result(nil, err, 0, false, c)
		return
	}
	err = validation.Validate(unit, validation.NotNil)
	if err != nil {
		msg.Result(nil, err, 2, false, c)
		return
	}
	err, data, success := services.SCreateUnit(unit)
	if !success {
		msg.Result(nil, err, 2, false, c)
		return
	}
	msg.Result(data, err, 0, true, c)
	return
}
