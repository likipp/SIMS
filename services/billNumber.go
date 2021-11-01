package services

import (
	"SIMS/models"
	"errors"
)

func SGenerateNumber(t string) (error, string) {
	number := models.GenerateNumber(t)
	if number == "" {
		return errors.New("出入库类型为空, 无法生成编码"), ""
	}
	return nil, number
}