package services

import (
	"SIMS/models"
)

func SCreateBrand(b *models.Brand) (err error, bR *models.Brand, success bool) {
	err, success = b.CreateBrand()
	return err, b, success
}
