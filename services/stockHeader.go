package services

import "SIMS/models"

func SStockHeader(sh models.StockHeader, sb []models.StockBody) (er error, success bool) {
	err, success := sh.StockHeaderLog(sb)
	return err, success
}

func SGetExBillDetail(number string) (err error, b models.ExBillDetail, success bool) {
	err, b, success = models.GetExBillDetail(number)
	return err, b, success
}
