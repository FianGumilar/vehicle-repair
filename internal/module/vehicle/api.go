package vehicle

import (
	"context"
	"time"

	"github.com/FianGumilar/vehicle-repair/domain"
	"github.com/FianGumilar/vehicle-repair/utils"
	"github.com/gofiber/fiber/v2"
)

type api struct {
	vehicleService domain.VehicleService
}

func NewApi(app *fiber.App, vehicleService domain.VehicleService) {
	api := api{
		vehicleService: vehicleService,
	}

	app.Get("/v1/vehicle-history", api.GetVehicleHistory)
	app.Post("/v1/vehicle-history", api.SaveVehicle)
}

func (a api) GetVehicleHistory(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	vin := ctx.Query("vin")
	if vin == "" {
		apiResponse := domain.ApiResponse{
			Code:    "99",
			Message: "Not Found",
		}
		utils.ResponseInterceptor(c, &apiResponse)
		ctx.Status(fiber.StatusBadRequest).JSON(apiResponse)
	}

	apiResponse := a.vehicleService.FindHistorical(c, vin)
	utils.ResponseInterceptor(c, &apiResponse)

	return ctx.Status(fiber.StatusOK).JSON(apiResponse)
}

func (a api) SaveVehicle(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	var req domain.VehicleHistoricalRequest

	if err := ctx.BodyParser(&req); err != nil {
		apiResponse := domain.ApiResponse{
			Code:    "01",
			Message: "Invalid Body",
		}
		utils.ResponseInterceptor(c, &apiResponse)

		return ctx.Status(fiber.StatusBadRequest).JSON(apiResponse)
	}

	apiResponse := a.vehicleService.StoreVehicleHistory(c, req)
	utils.ResponseInterceptor(c, &apiResponse)

	return ctx.Status(fiber.StatusOK).JSON(apiResponse)
}
