package config

import (
	"github.com/gin-gonic/gin"
	"aveonline/internal/modules/promotion"
	"aveonline/internal/modules/invoice"
	"aveonline/internal/modules/medicine"
	"github.com/gin-contrib/cors"
)
type Routes struct{
	app *Server
	handlerPromotion *promotion.PromotionHandler
	handlerInvoice *invoice.InvoiceHandler
	handlerMedicine *medicine.MedicineHandler
}
func NewRoutes(
	server *Server,
	promotion *promotion.PromotionHandler,
	invoice *invoice.InvoiceHandler,
	medicine *medicine.MedicineHandler,
) *Routes{
	return &Routes{app: server, handlerPromotion: promotion, handlerInvoice: invoice, handlerMedicine: medicine}
}
func(server *Routes) RegisterRouters(){
	server.app.Gin.Use(cors.Default())
	server.app.Gin.GET("/",func(c *gin.Context){
		c.String(200,"Systes Ready")
	})
	v1:= server.app.Gin.Group("/v1")
	{
		//swagger:route GET /promotion promotion promotionList
		// return all promotion
		// responses:
		//   200: []Promotion
		//   500: serverError
		v1.GET("/promotion",server.handlerPromotion.ListPromotions)

		//swagger:route POST /promotion promotion promotionCreate
		// create new promotion
		// responses:
		//   200: Promotion
		//   500: serverError
		v1.POST("/promotion",server.handlerPromotion.CreatePromotion)

		// swagger:route GET /v1/invoice invoice getInvoices
		//	return all invoices
		// responses:
		// 	200: []Invoice
		// 	500: error
		v1.GET("/invoice",server.handlerInvoice.ListInvoices)

		// swagger:route POST /v1/invoice/ invoice postInvoice
		//	create invoice
		// responses:
		// 	200: Invoice
		// 	500: error
		v1.POST("/invoice",server.handlerInvoice.CreateInvoice)

		
		v1.GET("/invoice/:startDate/:finishDate",server.handlerInvoice.GetPaymentForDate)

		// swagger:route GET /v1/medicine medicine getMedicines
		// return all medicines
		// responses:
		// 	200: []Medicine
		// 	500: error
		v1.GET("/medicine",server.handlerMedicine.ListMedicines)

		// swagger:route POST /v1/medicine medicine postMedicine
		//	create medicine
		// responses:
		// 	200: Medicine
		// 	500: error
		v1.POST("/medicine",server.handlerMedicine.CreateMedicine)

		v1.GET("/medicine/:id",server.handlerMedicine.FindByID)
	}
}