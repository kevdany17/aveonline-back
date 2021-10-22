package main

import (
	"aveonline/internal/config"
	"go.uber.org/dig"
	"aveonline/internal/modules/promotion"
	"aveonline/internal/modules/invoice"
	"aveonline/internal/modules/medicine"
)
// Prueba
func main(){
	container := LoadDependencies()

	err := container.Invoke(func(migrate *config.Migration, server *config.Server, routes *config.Routes) {
		migrate.InitMigrations()
		routes.RegisterRouters()
		server.InitServer()
	  })
	
	if err != nil {
		panic(err)
	}
	
}

func LoadDependencies() *dig.Container{
	container := dig.New()
	container.Provide(config.InitalizeDataBase)
	container.Provide(config.NewServer)
	container.Provide(config.NewRoutes)
	container.Provide(config.NewMigration)
	container.Provide(promotion.NewPromotionHandler)
	container.Provide(promotion.NewPromotionRepository)
	container.Provide(invoice.NewInvoiceHandler)
	container.Provide(invoice.NewInvoiceRepository)
	container.Provide(medicine.NewMedicineHandler)
	container.Provide(medicine.NewMedicineRepository)

	return container
}