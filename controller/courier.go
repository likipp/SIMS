package controller

import (
	"SIMS/models"
	"SIMS/services"
	"SIMS/utils/msg"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ACreateCourier(c *gin.Context) {
	var courier *models.Courier
	var _ = c.ShouldBind(&courier)
	fmt.Println(courier, "courier")
	//_ = c.ShouldBindBodyWith(&courier, binding.JSON).Error()
	err, data := services.SCreateCourier(courier)
	if err != nil {
		msg.FailWithMessage(msg.CreatedFail, c)
	} else {
		msg.SuccessWithData(data, c)
	}
}
