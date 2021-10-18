package controller

import (
	"SIMS/internal/gins"
	"SIMS/models"
	"SIMS/services"
	"SIMS/utils/msg"
	"github.com/gin-gonic/gin"
)

func CCreateCustom(c *gin.Context) {
	var custom *models.Custom
	err := gins.ParseJSON(c, &custom)
	if err != nil {
		msg.Result(nil, err.Error(), 0, false, c)
		return
	}
	err, data := services.SCreateCustom(custom)
	if err != nil {
		msg.Result(nil, err.Error(), 0, false, c)
		return
	}
	msg.Result(data, err.Error(), 0, true, c)
}
