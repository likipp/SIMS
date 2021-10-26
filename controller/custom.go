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
		msg.Result(nil, err, 0, false, c)
		return
	}
	err, data := services.SCreateCustom(custom)
	if err != nil {
		msg.Result(nil, err, 0, false, c)
		return
	}
	msg.Result(data, err, 0, true, c)
}

func GetCustomList(c *gin.Context) {
	var params models.CustomQueryParams
	err := gins.ParseQuery(c, &params)
	if err != nil {
		msg.Result(nil, err, 1, false, c)
		return
	}
	success, err, list, total := new(models.Custom).GetList(params)
	if !success {
		msg.Result(nil, err, 1, success, c)
		return
	}
	msg.ResultWithPageInfo(list, err, 1, success, total, params.Current, params.PageSize, c)
}


func CGetCustomList(c *gin.Context) {
	var customs []models.CustomSelect
	param := c.DefaultQuery("param", "")
	err, customs, success := services.SGetCustomList(param)
	if !success {
		msg.Result(nil, err, 2, false, c)
		return
	}
	msg.Result(customs, err, 1, true, c)
}