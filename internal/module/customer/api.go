package customer

import (
	"github.com/FianGumilar/vehicle-repair/domain"
	"github.com/gofiber/fiber/v2"
)

type api struct {
	customerService domain.CustomerService
}

func NewApi(app *fiber.App, customerService domain.CustomerService) {
	api := api{
		customerService: customerService,
	}

	app.Get("/welcome", api.Hi)
}

func (a api) Hi(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"code":    "200",
		"message": "Success",
	})
}
