package services

import "SIMS/models"

//func SInStock(s *models.Stock) (err error, sR *models.Stock, success bool) {
//	err, success = s.InStock()
//	return err, sR, success
//}
//
//func SExStock(s *models.Stock) (err error, sR *models.Stock, success bool) {
//	err, success = s.ExStock()
//	return err, sR, success
//}

// SGetStockList 获取即时库存
func SGetStockList(params models.StockQueryParams) (err error, sR []models.Stock, success bool) {
	err, list, success := models.GetStockList(params)
	if !success {
		return err, list, false
	}
	return err, list, true
}

func SChangeStock(s models.Stock, action string, qty int) (err error, success bool) {
	err, success = s.ChangeStock(action, qty)
	if !success {
		return err, false
	}
	return err, success
}

func SGetExStockList(params models.ExListQueryParams) (err error, list []models.ExStock, success bool) {
	err, list, success = models.GetExStockList(params)
	return err, list, success
}

func SGetInStockList(params models.InStockQueryParams) (err error, list []models.InStock, success bool) {
	err, list, success = models.GetInStockList(params)
	return err, list, success
}

func SGetInExStockList(params models.InExStockQueryParams) (err error, list []models.InExStock, success bool) {
	err, list, success = models.GetInExStockList(params)
	return err, list, success
}
