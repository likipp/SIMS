package services

import "SIMS/models"

func SStockHeader(sh models.BillHeader, sb []models.BillEntry) (err error, success bool) {
	err, success = sh.BillLog(sb)
	return err, success
}

func SDeleteBillByID(id int) (err error, success bool) {
	err, success = models.DeleteBillByID(id)
	return err, success
}

func SGetExBillDetail(number string) (err error, b models.ExBillDetail, success bool) {
	err, b, success = models.GetExBillDetail(number)
	return err, b, success
}

func SGetInBillDetail(number string) (err error, b models.InBillDetail, success bool) {
	err, b, success = models.GetInBillDetail(number)
	return err, b, success
}
