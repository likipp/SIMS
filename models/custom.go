package models

type Custom struct {
	BaseModel
	CName      string    `json:"c_name"`
	CNumber    string    `json:"c_number"`
	Phone      string    `json:"phone"`
	Address    string    `json:"address"`
}