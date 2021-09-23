package promotion

import (
	"gorm.io/gorm"
	"log"
)

type PromotionRepository struct{
	db *gorm.DB
}
func NewPromotionRepository(connection *gorm.DB) *PromotionRepository{
	return &PromotionRepository{db: connection}
}
func(repo *PromotionRepository) List() (*[]Promotion, error) {
	var data []Promotion
	result := repo.db.Model(&Promotion{}).Find(&data)
	if result.Error != nil{
		log.Fatal(result.Error)
		return nil, gorm.ErrRecordNotFound
	}
	return &data, nil
}
func(repo *PromotionRepository) Create(data Promotion) (*Promotion, error) {
	result := repo.db.Create(&data)
	if result.Error != nil{
		log.Fatal(result.Error)
		return nil, gorm.ErrRecordNotFound
	}
	return &data, nil
}
func(repo *PromotionRepository) GetPromotionForDate(startDate string, finishDate string) (*[]Promotion, error) {
	var data []Promotion
	result := repo.db.Model(&Promotion{}).Where("start_date = ? OR finish_date = ?",startDate, finishDate).Find(&data)
	if result.Error != nil{
		log.Fatal(result.Error)
		return nil, gorm.ErrRecordNotFound
	}
	return &data, nil
}