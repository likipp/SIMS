package services

import (
	"SIMS/models"
	"SIMS/utils/msg"
)

func SCreateCustomLevel(c *models.CustomLevel) (err error, cR *models.CustomLevel) {
	err = c.CreateCustomLevel()
	if err != nil {
		return msg.CreatedFail, c
	}
	return msg.CreatedSuccess, c
}

func SGetCustomLevelList() (err error, crs []*models.CustomLevel) {
	//
	return err, crs
}

func SUpdateCustomLevel(c *models.CustomLevel, id int) (err error) {
	err = c.UpdateCustomLevel(id)
	return err
}

func SDeleteCustomLevel(c *models.CustomLevel) (err error) {
	err = c.DeleteCustomLevel()
	return err
}
