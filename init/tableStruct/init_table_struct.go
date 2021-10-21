package tableStruct

import (
	"SIMS/models"
	"gorm.io/gorm"
)

func InitTableStruct(db *gorm.DB) {
	_ = db.AutoMigrate(
		models.WareHouse{},
		models.Brand{},
		models.Products{},
		models.Courier{},
		models.CustomLevel{},
		models.Custom{},
		models.Stock{},
	)
}
