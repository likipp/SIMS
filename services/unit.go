package services

import "SIMS/models"

func SCreateUnit(u *models.Unit) (err error, ur *models.Unit, success bool) {
	err = u.CreateUnit()
	if err != nil {
		return err, u, false
	}
	return err, u, true
}

func SGetUnitSelectList(param string) (error, []models.UnitSelect, bool) {
	err, list, success := models.GetUnitSelectList(param)
	return err, list, success
}
