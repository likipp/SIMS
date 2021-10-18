package models

type Custom struct {
	BaseModel
	CName    string `json:"c_name"    gorm:"comment:'客户名称'"`
	CNumber  string `json:"c_number"  gorm:"comment:'顾客编码'"`
	Phone    string `json:"phone"     gorm:"comment:'手机号码'"`
	Address  string `json:"address"   gorm:"comment:'收件件地址'"`
	Discount int    `json:"discount"  gorm:"comment:'折扣'"`
	Level    int    `json:"level"     gorm:"level:'客户等级'"`
}
