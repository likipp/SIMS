package services

import (
	"SIMS/models"
)

type Courier struct {
	CName   string `json:"c_name"`
	CNumber string `json:"c_number"`
}

func SCreateCourier(c *models.Courier) (err error, cR *models.Courier) {
	err = c.CreateCourier()
	if err != nil {
		return err, c
	}
	return err, c
}

func SUpdateCourier(c *models.Courier, id int) (err error, cR *models.Courier) {
	err = c.UpdateCourier(id)
	if err != nil {
		return err, c
	}
	return err, c
}

func SDeleteCourier(c *models.Courier) (err error) {
	err = c.DeleteCourier()
	return err
}
