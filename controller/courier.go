package controller

import (
	"SIMS/models"
	"SIMS/services"
	"SIMS/utils"
	"SIMS/utils/msg"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ACreateCourier(c *gin.Context) {
	var courier *models.Courier
	var _ = c.ShouldBind(&courier)
	//_ = c.ShouldBindBodyWith(&courier, binding.JSON).Error()
	err, data := services.SCreateCourier(courier)
	if err != nil {
		msg.FailWithMessage(msg.CreatedFail, c)
	} else {
		msg.SuccessWithData(data, c)
	}
}

func AUpdateCourier(c *gin.Context) {
	fmt.Println(c.Request.Method, "Request")
	var r *models.Courier
	_ = c.ShouldBindJSON(&r)
	id := c.Param("id")
	err := services.SUpdateCourier(r, utils.StringConvInt(id))
	if err != nil {
		msg.FailWithMessage(msg.UpdatedFail, c)
	} else {
		msg.SuccessWithMessage(msg.UpdatedSuccess, c)
	}
}

func ADeleteCourier(c *gin.Context) {
	var r models.Courier
	id := c.Param("id")
	r.ID = utils.StringConvInt(id)
	err := services.SDeleteCourier(&r)
	if err != nil {
		msg.Fail(c)
	} else {
		msg.Success(c)
	}
}
