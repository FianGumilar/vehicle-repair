package customer

import (
	"context"
	"time"

	"github.com/FianGumilar/vehicle-repair/domain"
	"github.com/FianGumilar/vehicle-repair/utils"
	"github.com/gofiber/fiber/v2"
)

type api struct {
	customerService domain.CustomerService
}

func NewApi(app *fiber.App, customerService domain.CustomerService) {
	api := api{
		customerService: customerService,
	}

	app.Get("/v1/customers", api.AllCustomers)
	app.Post("v1/customers", api.SaveCustomer)

}

func (a api) AllCustomers(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	apiResponse := a.customerService.All(c)
	utils.ResponseInterceptor(c, &apiResponse)

	return ctx.Status(fiber.StatusOK).JSON(apiResponse)
}

func (a api) SaveCustomer(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	var customerData domain.CustomerData

	if err := ctx.BodyParser(&customerData); err != nil {
		apiResponse := domain.ApiResponse{
			Code:    "02",
			Message: "Invalid Parameter",
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(apiResponse)
	}

	apiResponse := a.customerService.Save(c, customerData)
	utils.ResponseInterceptor(c, &apiResponse)

	return ctx.Status(fiber.StatusOK).JSON(apiResponse)
}
