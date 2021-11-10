package controller

import (
	"SIMS/internal/gins"
	"SIMS/models"
	"SIMS/services"
	"SIMS/utils/msg"
	"github.com/gin-gonic/gin"
)

func CCreateCustomLevel(c *gin.Context) {
	var customLevel *models.CustomLevel
	err := c.BindJSON(&customLevel)
	//err, ok := err.(validator.ValidationErrors)
	//fmt.Println(err, ok)
	//if !ok {
	//	msg.Result(nil, err, 0, false, c)
	//	return
	//}
	err, data := services.SCreateCustomLevel(customLevel)
	if err != nil {
		msg.Result(nil, err, 1, false, c)
		return
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
		msg.Result(nil, err, 0, success, c)
		return
	}
	msg.ResultWithPageInfo(list, err, 1, success, total, params.Current, params.PageSize, c)
}

func CGetCustomLevelSelectList(c *gin.Context) {
	var cls []models.CustomLevelSelect
	param := c.DefaultQuery("param", "")
	if param == "{}" {
		param = ""
	}
	err, cls, success := services.SGetCustomLevelSelectList(param)
	if !success {
		msg.Result(nil, err, 2, false, c)
		return
	}
	msg.Result(cls, err, 1, true, c)
}
