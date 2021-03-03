package models

import "time"

type BaseModel struct {
	CreatedAt *time.Time `gorm:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"updated_at" json:"updatedAt"`
	DeleteAt  *time.Time `gorm:"delete_at"  json:"deleteAt"`
	CreateBy  string     `gorm:"create_by" json:"createBy"`
	UpdateBy  string     `gorm:"update_by" json:"updateBy"`
	DeleteBy  string     `gorm:"delete_by" json:"deleteBy"`
	ID        int        `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
}