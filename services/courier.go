package services

import "SIMS/models"

type Courier struct {
	CName   string `json:"c_name"`
	CNumber string `json:"c_number"`
}


func SCreateCourier(c *models.Courier) (err error, cR *Courier) {
	err = c.CreateCourier()
	return err, nil
}