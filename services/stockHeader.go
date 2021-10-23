package services

import "SIMS/models"

func SStockHeader(sh models.StockHeader, sb []models.StockBody) (er error, success bool) {
	err, success := sh.StockHeaderLog(sb)
	return err, success
}
