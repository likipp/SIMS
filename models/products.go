package models

type Products struct {
	BaseModel
	PName   string `json:"p_name" gorm:"comment:'产品名称'"`
	PNumber string `gorm:"index not null; comment:'产品编号'" json:"p_number"`
	Spec    string `json:"p_spec" gorm:"comment:'规格型号'"`
	Price   string `json:"p_price" gorm:"comment:'单价'"`
	Picture string `json:"picture" gorm:"default:'/favicon.ico'; comment:'图片'"`
}
