package models

type Products struct {
	BaseModel
	PName   string `json:"p_name"`
	PNumber string `gorm:"index not null" json:"p_number"`
	Spec    string `json:"p_spec"`
	Price   string `json:"p_price"`
	Picture string `json:"picture" gorm:"default:'/favicon.ico'"`
}
