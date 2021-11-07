package services

import (
	"SIMS/models"
)

func SCreateBrand(b *models.Brand) (err error, bR *models.Brand, success bool) {
	err, success = b.CreateBrand()
	return err, b, success
}

func SGetBrandSelectList(param string) (error, []models.BrandSelect, bool) {
	err, list, success := models.GetBrandSelectList(param)
	return err, list, success
}

func SGetBrandTree() (error, []models.BrandTree, bool) {
	err, list, success := models.GetBrandTree()
	return err, list, success
}
