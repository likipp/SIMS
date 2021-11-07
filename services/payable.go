package services

import "SIMS/models"

func SCreatePayBill(p *models.Payable) (err error, success bool) {
	err, success = p.CreatePayBill()
	return err, success
}
