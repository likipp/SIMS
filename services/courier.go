package services

import (
	"SIMS/models"
	"SIMS/utils/msg"
)

type Courier struct {
	CName   string `json:"c_name"`
	CNumber string `json:"c_number"`
}

func SCreateCourier(c *models.Courier) (err error, cR *models.Courier) {
	err = c.CreateCourier()
	if err != nil {
		return msg.CreatedFail, c
	}
	return msg.CreatedSuccess, c
}

func SUpdateCourier(c *models.Courier, id int) (err error) {
	err = c.UpdateCourier(id)
	if err != nil {
		return err
	}
	return err
}

func SDeleteCourier(c *models.Courier) (err error) {
	err = c.DeleteCourier()
	return err
}
