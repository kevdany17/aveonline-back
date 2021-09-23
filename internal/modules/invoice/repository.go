package invoice

import (
	"gorm.io/gorm"
	"log"
)

type InvoiceRepository struct{
	db *gorm.DB
}
func NewInvoiceRepository(connection *gorm.DB) *InvoiceRepository{
	return &InvoiceRepository{db: connection}
}
func(repo *InvoiceRepository) List() (*[]Invoice, error) {
	var data []Invoice
	result := repo.db.Preload("Promotion").Preload("Medicines").Model(&Invoice{}).Find(&data)
	if result.Error != nil{
		log.Fatal(result.Error)
		return nil, gorm.ErrRecordNotFound
	}
	return &data, nil
}
func(repo *InvoiceRepository) Create(data Invoice) (*Invoice, error) {
	result := repo.db.Create(&data)
	if result.Error != nil{
		log.Fatal(result.Error)
		return nil, gorm.ErrRecordNotFound
	}
	return &data, nil
}
func(repo *InvoiceRepository) GetPaymentsForDate(startDate string, finisDate string) (*[]PaymentForDate, error) {
	var data []PaymentForDate
	result := repo.db.Raw(`SELECT creation_date AS date, sum(total_pay) as total  from invoices
							WHERE creation_date BETWEEN ? AND ?
							GROUP BY creation_date;`, startDate, finisDate).Scan(&data)
	if result.Error != nil{
		log.Fatal(result.Error)
		return nil, gorm.ErrRecordNotFound
	}
	return &data, nil
}