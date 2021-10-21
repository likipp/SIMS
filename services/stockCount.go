package services

import "SIMS/models"

func SStockCount(s *models.StockCount) (err error, success bool) {
	err, success = s.StockCount()
	return err, success
}