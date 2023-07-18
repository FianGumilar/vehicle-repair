package main

import (
	"github.com/FianGumilar/vehicle-repair/internal/component"
	"github.com/FianGumilar/vehicle-repair/internal/config"
	"github.com/FianGumilar/vehicle-repair/internal/module/customer"
	"github.com/gofiber/fiber/v2"
)

func main() {
	conf := config.Get()

	dbConnection := component.GetDatabaseConnection(conf)

	customerRepository := customer.NewRepository(dbConnection)

	customerService := customer.NewService(customerRepository)

	app := fiber.New()
	customer.NewApi(app, customerService)

	app.Listen(conf.Srv.Host + ":" + conf.Srv.Port)
}
