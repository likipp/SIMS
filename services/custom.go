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
