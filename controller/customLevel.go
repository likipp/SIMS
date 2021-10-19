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
		msg.Result(nil, err, 0, false, c)
		return
	}
	err, data := services.SCreateCustomLevel(customLevel)
	if err != nil {
		msg.Result(nil, err, 0, false, c)
		//return
	}
	msg.Result(data, err, 0, true, c)
}

func GetCustomLevelList(c *gin.Context) {
	var params models.CustomLevelQueryParams
	err := gins.ParseQuery(c, &params)
	if err != nil {
		msg.Result(nil, err, 1, false, c)
		return
	}
	success, err, list, total := new(models.CustomLevel).GetList(params)
	if !success {
		msg.Result(nil, err, 1, success, c)
		return
	}
	msg.ResultWithPageInfo(list, err, 1, success, total, params.Current, params.PageSize, c)
}
