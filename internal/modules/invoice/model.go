package invoice

import (
	"aveonline/internal/modules/promotion"
	"aveonline/internal/modules/medicine"
)
//swagger:model Invoice
type Invoice struct{
	ID				uint					`gorm:"primaryKey" json:"id"`	
	CreationDate 	string  				`json:"creationDate"`    
	TotalPay 		float64					`json:"totalPay"`
	PromotionID		uint					`json:"promotionId"`
	Promotion		promotion.Promotion		`json:"promotion"`
	Medicines		[]medicine.Medicine		`gorm:"many2many:invoice_medicines; json:medicines"`
}

//swagger:model PaymentForDate
type PaymentForDate struct{
	Date 	string 		`json:"name"`
	Total	float64		`json:"value"`
}