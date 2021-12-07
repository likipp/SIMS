package services

import "SIMS/models"

func SCreatePayBill(p *models.Payable) (err error, success bool) {
	err, success = p.CreatePayBill()
	return err, success
}

func SGetPayList(param int) (error, []models.PayList, bool) {
	err, data, success := models.GetPayList(param)
	return err, data, success
}

func SGetPayPie() (error, []models.PayPie, bool) {
	err, data, success := models.GetPayPie()
	return err, data, success
}

func SGetExColumn() (error, []models.ExColumn, bool) {
	err, data, success := models.GetExColumn()
	return err, data, success
}

func SGetProductSale() (error, []models.ProductSale, bool) {
	err, data, success := models.GetProductSale()
	return err, data, success
}