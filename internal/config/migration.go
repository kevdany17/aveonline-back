package config

import (
	"aveonline/internal/modules/promotion"
	"aveonline/internal/modules/invoice"
	"aveonline/internal/modules/medicine"
	"gorm.io/gorm"
)

type Migration struct{
	db *gorm.DB
}
func NewMigration(connection *gorm.DB) *Migration{
	return &Migration{db: connection}
}
func(migrate *Migration) InitMigrations(){
	migrate.db.AutoMigrate(&promotion.Promotion{})
	migrate.db.AutoMigrate(&invoice.Invoice{})
	migrate.db.AutoMigrate(&medicine.Medicine{})
}