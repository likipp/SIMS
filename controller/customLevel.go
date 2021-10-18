package controller

import (
	"SIMS/internal/gins"
	"SIMS/models"
	"SIMS/services"
	"SIMS/utils/msg"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CCreateCustomLevel(c *gin.Context) {
	var customLevel *models.CustomLevel
	err := c.BindJSON(&customLevel)
	err, ok := err.(validator.ValidationErrors)
	if !ok {
		msg.Result(nil, err.Error(), 0, false, c)
		return
	}
	err, data := services.SCreateCustomLevel(customLevel)
	if err != nil {
		msg.Result(nil, err.Error(), 0, false, c)
		//return
	}
	msg.Result(data, err.Error(), 0, true, c)
}

func GetList(c *gin.Context) {
	var params models.QueryParams
	err := gins.ParseQuery(c, &params)
	if err != nil {
		msg.Result(nil, err.Error(), 1, false, c)
	}
	err, list, total := new(models.CustomLevel).GetCustomLevel(params)
	if err != nil {
		msg.Result(nil, err.Error(), 1, false, c)
	} else {
		msg.ResultWithPageInfo(list, err.Error(), 1, true, total, params.Current, params.PageSize, c)
	}
}
