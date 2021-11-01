package controller

import (
	"SIMS/services"
	"SIMS/utils/msg"
	"github.com/gin-gonic/gin"
)

func CGenerateNumber(c *gin.Context) {
	t := c.Query("type")
	err, number := services.SGenerateNumber(t)
	if err != nil {
		msg.Result(nil, err, 1, false, c)
		return
	}
	msg.Result(number, msg.GetSuccess, 1, true, c)
	return
}