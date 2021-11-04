package services

import (
	"SIMS/models"
)

func SCreateProduct(p *models.Products) (err error, pR *models.Products, success bool) {
	err, success = p.CreateProducts()
	return err, pR, success
}

func SGetProductsSelectList(param string) (err error, pr []models.ProductsSelect, success bool) {
	err, pr, success = models.GetProductsSelectList(param)
	return err, pr, success
}
