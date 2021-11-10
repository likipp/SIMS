package services

import (
	"SIMS/models"
)

func SCreateCustom(c *models.Custom) (err error, success bool) {
	err, success = c.CreateCustom()
	return err, success
}

func SGetCustomByID(id int) (err error, c models.Custom, success bool) {
	success, err, c = models.GetCustomByID(id)
	return err, c, success
}

func SGetCustomListWithQuery(param string) (err error, pr []models.CustomSelect, success bool) {
	err, pr, success = models.GetCustomsListWithQuery(param)
	return err, pr, success
}
