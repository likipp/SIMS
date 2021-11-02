package router

import (
	"SIMS/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	baseRouter := r.Group("/api/v1/base")
	{
		baseRouter.POST("login/", controller.Login)
		baseRouter.GET("currentUser/", controller.GetCurrentUser)
		// 供应商路由
		baseRouter.POST("couriers", controller.ACreateCourier)
		baseRouter.PATCH("couriers/:id", controller.AUpdateCourier)
		baseRouter.DELETE("couriers/:id", controller.ADeleteCourier)

		// 客户等级路由
		baseRouter.POST("custom-level", controller.CCreateCustomLevel)
		baseRouter.GET("custom-level/", controller.GetCustomLevelList)

		// 客户路由
		baseRouter.POST("custom", controller.CCreateCustom)
		baseRouter.GET("custom/select", controller.CGetCustomListWithQuery)
		baseRouter.GET("custom/", controller.CGetCustomList)

		// 品牌路由
		baseRouter.POST("brand", controller.CCreateBrand)

		// 仓库路由  CWareHouseList
		baseRouter.POST("warehouse/", controller.CCreateWareHouse)
		baseRouter.GET("warehouse/", controller.CWareHouseList)

		// 产品路由
		baseRouter.POST("product/", controller.CCreateProduct)
		baseRouter.GET("product/", controller.CGetProductList)

		// 库存路由
		//baseRouter.POST("stock/", controller.CStockCount)
		// 获取单据编号
		baseRouter.GET("generate-number/", controller.CGenerateNumber)
		// 创建出入库单据
		baseRouter.POST("stock/", controller.CStockHeader)
		baseRouter.GET("stock/", controller.CGetStockList)
		baseRouter.GET("stock/ex/", controller.CGetExStockList)
		baseRouter.GET("stock/in", controller.CGetInStockList)
		baseRouter.POST("stock/:id/", controller.CSChangeStock)
		baseRouter.GET("ex-bill/", controller.CGetExBillDetail)
	}
	return r
}
