package services

import "SIMS/models"

func SCreateWareHouse(w *models.WareHouse) (err error, wR *models.WareHouse, success bool) {
	err, success = w.CreateWareHouse()
	return err, wR, success
}

func SGetWareHouseSelectList(param string) (err error, pr []models.WareHouseSelect, success bool) {
	err, pr, success = models.GetWareHouseSelectList(param)
	return err, pr, success
}

//
//func SGetWareHouseList(param string) (err error, pr []models.WareHouseSelect, success bool) {
//	err, pr, success = models.GetWareHouseList(param)
//	return err, pr, success
//}
