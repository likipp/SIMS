package router

import (
	"SIMS/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	baseRouter := r.Group("/api/v1/base")
	{
		baseRouter.POST("couriers", controller.ACreateCourier)
		baseRouter.PATCH("couriers/:id", controller.AUpdateCourier)
		baseRouter.DELETE("couriers/:id", controller.ADeleteCourier)
		baseRouter.POST("custom-level", controller.CCreateCustomLevel)
		baseRouter.GET("custom-level", controller.GetList)
	}
	return r
}
