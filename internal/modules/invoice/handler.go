package invoice

import (
	"github.com/gin-gonic/gin"
	"time"
)

type InvoiceHandler struct{
	repository *InvoiceRepository
}
func NewInvoiceHandler(repo *InvoiceRepository) *InvoiceHandler{
	return &InvoiceHandler{repository: repo}
}
func(handler *InvoiceHandler) ListInvoices(c *gin.Context){
	data, err := handler.repository.List()
	if err != nil{
		c.JSON(500,gin.H{
			"error": err,
		})
	}
	c.JSON(200,&data)
}
func(handler *InvoiceHandler) CreateInvoice(c *gin.Context){
	var invoice Invoice
	err := c.ShouldBindJSON(&invoice)
	if err != nil{
		c.JSON(500,gin.H{
			"error": "Incorrect Format",
		})
		return
	}
	today := time.Now()
	invoice.CreationDate = today.Format("2006-01-02")
	data, err := handler.repository.Create(invoice)
	if err != nil{
	 	c.JSON(500,gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200,&data)
}

func(handler *InvoiceHandler) GetPaymentForDate(c *gin.Context){
	startDate := c.Param("startDate")
	finishDate := c.Param("finishDate")
	data, err := handler.repository.GetPaymentsForDate(startDate,finishDate)
	if err != nil{
		c.JSON(500,gin.H{
			"error": err,
		})
	}
	c.JSON(200,&data)
}