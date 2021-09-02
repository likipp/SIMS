package models

import "time"

// 出货表

type ShipmentHeader struct {
	SNumber  string     `json:"s_number"`
	STime    *time.Time `json:"s_time"`
	SCustom  string     `json:"s_custom"`
	SAddress string     `json:"s_address"`
	Courier  string     `json:"courier"`
}

type ShipmentBody struct {
	SNumber   string  `json:"s_number"`
	PName     string  `json:"p_name"`
	PNumber   string  `json:"p_number"`
	Amount    int     `json:"amount"`
	UnitPrice float32 `json:"unit_price"`
	Total     float32 `json:"total"`
}
