package router

import (
	"SIMS/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	baseRouter := r.Group("/api/v1/base")
	{
		// 供应商路由
		baseRouter.POST("couriers", controller.ACreateCourier)
		baseRouter.PATCH("couriers/:id", controller.AUpdateCourier)
		baseRouter.DELETE("couriers/:id", controller.ADeleteCourier)

		// 客户等级路由
		baseRouter.POST("custom-level", controller.CCreateCustomLevel)
		baseRouter.GET("custom-level/", controller.GetCustomLevelList)

		// 客户路由
		baseRouter.POST("custom", controller.CCreateCustom)
		baseRouter.GET("custom/", controller.GetCustomList)

		// 品牌路由
		baseRouter.POST("brand", controller.CCreateBrand)

		// 仓库路由
		baseRouter.POST("warehouse/", controller.CCreateWareHouse)

		//
		baseRouter.POST("product/", controller.CCreateProduct)
	}
	return r
}
