package services

import (
	"SIMS/models"
)

func SCreateCustom(c *models.Custom) (err error, cR *models.Custom) {
	err = c.CreateCustom()
	if err != nil {
		return err, c
	}
	return nil, c
}

func SGetCustomByID(id int) (err error, c models.Custom, success bool) {
	success, err, c = models.GetCustomByID(id)
	return err, c, success
}

func SGetCustomListWithQuery(param string) (err error, pr []models.CustomSelect, success bool) {
	err, pr, success = models.GetCustomsListWithQuery(param)
	return err, pr, success
}
