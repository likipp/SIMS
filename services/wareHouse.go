package services

import "SIMS/models"

func SCreateWareHouse(w *models.WareHouse) (err error, wR *models.WareHouse, success bool) {
	err, success = w.CreateWareHouse()
	return err, wR, success
}