package router

import (
	"SIMS/controller"
	"SIMS/middleware"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func InitRouter() {
	r := gin.Default()
	//r.StaticFS("/images", http.Dir("./static/images"))
	//r.StaticFile("/favicon.ico", "./static/images/default.jpg")
	r.POST("/api/v1/base/login/", controller.Login)
	r.Use(middleware.JWTAuth()).Use(middleware.CorsMiddlewares())
	baseRouter := r.Group("/api/v1/base")
	{
		//baseRouter.POST("login/", controller.Login)
		baseRouter.POST("logout", controller.Logout)
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
		baseRouter.GET("customQuery", controller.CGetCustomListWithQuery)
		baseRouter.GET("custom/", controller.CGetCustomList)
		//baseRouter.GET("custom/", controller.CGetCustomByID)

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
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("服务关闭中...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("服务关闭完成")
}
