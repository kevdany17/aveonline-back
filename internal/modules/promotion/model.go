package promotion

import (
	"time"
)

//swagger:model Promotion
type Promotion struct{
	ID				uint		`gorm:"primaryKey" json:"id"`	
	Description 	string  	`json:"description"`    
	Percentage 		float64		`json:"percentaje"`
	StartDate		string		`json:"startDate"`
	FinishDate		string		`json:"finishDate"`
	CreatedAt 		*time.Time	`json:"-" `
  	UpdatedAt 		*time.Time	`json:"-"`
}