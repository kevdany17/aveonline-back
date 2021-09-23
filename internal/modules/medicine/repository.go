package medicine

import (
	"gorm.io/gorm"
	"log"
)

type MedicineRepository struct{
	db *gorm.DB
}
func NewMedicineRepository(connection *gorm.DB) *MedicineRepository{
	return &MedicineRepository{db: connection}
}
func(repo *MedicineRepository) List() (*[]Medicine, error) {
	var data []Medicine
	result := repo.db.Model(&Medicine{}).Find(&data)
	if result.Error != nil{
		log.Fatal(result.Error)
		return nil, gorm.ErrRecordNotFound
	}
	return &data, nil
}
func(repo *MedicineRepository) Create(data Medicine) (*Medicine, error) {
	result := repo.db.Create(&data)
	if result.Error != nil{
		log.Fatal(result.Error)
		return nil, gorm.ErrRecordNotFound
	}
	return &data, nil
}
func(repo *MedicineRepository) FindByID(id uint) (*Medicine, error) {
	var data Medicine
	result := repo.db.Where("id = ?", id).First(&data)
	if result.Error != nil{
		log.Fatal(result.Error)
		return nil, gorm.ErrRecordNotFound
	}
	return &data, nil
}