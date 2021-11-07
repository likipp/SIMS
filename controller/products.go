package controller

import (
	"SIMS/internal/gins"
	"SIMS/models"
	"SIMS/services"
	"SIMS/utils"
	"SIMS/utils/msg"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func CCreateProduct(c *gin.Context) {
	var product *models.Products
	err := gins.ParseJSON(c, &product)
	if err != nil {
		msg.Result(nil, err, 2, false, c)
		return
	}
	err = validation.Validate(product, validation.NotNil)
	if err != nil {
		msg.Result(nil, err, 2, false, c)
		return
	}
	err, data, success := services.SCreateProduct(product)
	if success {
		msg.Result(data, err, 0, true, c)
		return
	}
	msg.Result(nil, err, 1, false, c)
}

func CGetProductsSelectList(c *gin.Context) {
	var products []models.ProductsSelect
	param := c.DefaultQuery("param", "")
	if param == "{}" {
		param = ""
	}
	err, products, success := services.SGetProductsSelectList(param)
	if !success {
		msg.Result(nil, err, 2, false, c)
		return
	}
	msg.Result(products, err, 1, true, c)
}

func CGetProductsList(c *gin.Context) {
	var params models.ProductQueryParams
	err := gins.ParseQuery(c, &params)
	if err != nil {
		msg.Result(nil, err, 1, false, c)
		return
	}
	err, list, total, success := new(models.Products).GetProductsList(params)
	if !success {
		msg.Result(nil, err, 1, success, c)
		return
	}
	msg.ResultWithPageInfo(list, err, 0, success, total, params.Current, params.PageSize, c)
}

func CUpdateProduct(c *gin.Context) {
	var product models.Products
	err := gins.ParseJSON(c, &product)
	if err != nil {
		msg.Result(nil, err, 1, false, c)
		return
	}
	// 获取到UUID, 只有uuid有值时才能更新成功
	id := c.Param("id")
	product.ID = utils.StringConvInt(id)
	err, success := product.UpdateProduct()
	if !success {
		msg.Result(nil, err, 1, success, c)
		return
	}
	msg.Result(nil, err, 0, true, c)
}
