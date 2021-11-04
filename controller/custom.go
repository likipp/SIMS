package controller

import (
	"SIMS/internal/gins"
	"SIMS/models"
	"SIMS/services"
	"SIMS/utils"
	"SIMS/utils/msg"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CCreateCustom(c *gin.Context) {
	var custom *models.Custom
	err := gins.ParseJSON(c, &custom)
	fmt.Println(custom, "custom")
	if err != nil {
		msg.Result(nil, err, 0, false, c)
		return
	}
	err, data := services.SCreateCustom(custom)
	if err != nil {
		msg.Result(nil, err, 1, false, c)
		return
	}
	msg.Result(data, err, 0, true, c)
	return
}

func CGetCustomList(c *gin.Context) {
	var params models.CustomQueryParams
	err := gins.ParseQuery(c, &params)
	if err != nil {
		msg.Result(nil, err, 1, false, c)
		return
	}
	id := utils.StringConvInt(c.Query("id"))
	if id != 0 {
		err, custom, success := services.SGetCustomByID(id)
		if !success {
			msg.Result(nil, err, 2, false, c)
			return
		}
		msg.Result(custom, err, 0, true, c)
		return
	}
	success, err, list, total := new(models.Custom).GetList(params)
	if !success {
		msg.Result(nil, err, 1, success, c)
		return
	}
	msg.ResultWithPageInfo(list, err, 1, success, total, params.Current, params.PageSize, c)
}

func CGetCustomListWithQuery(c *gin.Context) {
	var customs []models.CustomSelect
	param := c.DefaultQuery("param", "")
	err, customs, success := services.SGetCustomListWithQuery(param)
	if !success {
		msg.Result(nil, err, 2, false, c)
		return
	}
	msg.Result(customs, err, 1, true, c)
	return
}
