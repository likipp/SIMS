package models

import "time"

// 进货表

type WarehousingHeader struct {
	WNumber string     `json:"s_number"`
	WTime   *time.Time `json:"s_time"`
	Courier string     `json:"courier"`
}

type WarehousingBody struct {
	WNumber   string  `json:"s_number"`
	PName     string  `json:"p_name"`
	PNumber   string  `json:"p_number"`
	Amount    int     `json:"amount"`
	UnitPrice float32 `json:"unit_price"`
	Total     float32 `json:"total"`
}
