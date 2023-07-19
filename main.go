package main

import (
	"github.com/FianGumilar/vehicle-repair/internal/component"
	"github.com/FianGumilar/vehicle-repair/internal/config"
	"github.com/FianGumilar/vehicle-repair/internal/module/customer"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	conf := config.Get()

	dbConnection := component.GetDatabaseConnection(conf)

	customerRepository := customer.NewRepository(dbConnection)

	customerService := customer.NewService(customerRepository)

	app := fiber.New()
	app.Use(requestid.New())
	app.Use(logger.New())
	customer.NewApi(app, customerService)

	app.Listen(conf.Srv.Host + ":" + conf.Srv.Port)
}
