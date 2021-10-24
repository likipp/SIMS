package models

type BillBase struct {
	BaseModel
	Name       string `json:"name" gorm:"comment:'名称'"`
	TemplateID string `json:"template_id" gorm:"comment:'简写'"`
}
