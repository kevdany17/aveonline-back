package medicine


//swagger:model Medicine
type Medicine struct{
	ID				uint					`gorm:"primaryKey" json:"id"`	
	Name 			string  				`json:"name"`    
	Price 			float64					`json:"price"`
	Location		string					`json:"location"`
}