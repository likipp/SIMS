package services

import (
	"SIMS/models"
)

func SLogin(l models.Login) (err error, success bool) {
	err, success = models.UserLogin(&l)
	return err, success
}
