package medicine

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type MedicineHandler struct{
	repository *MedicineRepository
}
func NewMedicineHandler(repo *MedicineRepository) *MedicineHandler{
	return &MedicineHandler{repository: repo}
}
func(handler *MedicineHandler) ListMedicines(c *gin.Context){
	data, err := handler.repository.List()
	if err != nil{
		c.JSON(500,gin.H{
			"error": err,
		})
	}
	c.JSON(200,&data)
}
func(handler *MedicineHandler) CreateMedicine(c *gin.Context){
	var medicine Medicine
	err := c.ShouldBindJSON(&medicine)
	if err != nil{
		c.JSON(500,gin.H{
			"error": "Incorrect Format",
		})
		return
	}
	if medicine.Name == ""{
		c.JSON(500,gin.H{
			"error": "Empty Name",
		})
		return
	}
	if len(medicine.Name) > 80{
		c.JSON(500,gin.H{
			"error": "Name max 80 characters",
		})
		return
	}
	if medicine.Price < 0.1{
		c.JSON(500,gin.H{
			"error": "Null Price",
		})
		return
	}
	if medicine.Location == ""{
		c.JSON(500,gin.H{
			"error": "Empty Location",
		})
		return
	}
	if len(medicine.Location) > 120{
		c.JSON(500,gin.H{
			"error": "Location max 120 characters",
		})
		return
	}
	data, err := handler.repository.Create(medicine)
	if err != nil{
	 	c.JSON(500,gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200,&data)
}
func(handler *MedicineHandler) FindByID(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := handler.repository.FindByID(uint(id))
	if err != nil{
	 	c.JSON(500,gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200,&data)
}