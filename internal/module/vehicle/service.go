package vehicle

import (
	"context"
	"time"

	"github.com/FianGumilar/vehicle-repair/domain"
)

type service struct {
	vehicleRepository domain.VehicleRepository
	historyRepository domain.HistoryRepository
}

func NewService(vehicleRepository domain.VehicleRepository, historyRepository domain.HistoryRepository) domain.VehicleService {
	return &service{vehicleRepository: vehicleRepository, historyRepository: historyRepository}
}

// FindHistorical implements domain.VehicleService.
func (s service) FindHistorical(ctx context.Context, vin string) domain.ApiResponse {
	vehicle, err := s.vehicleRepository.FindByVin(ctx, vin)
	if err != nil {
		return domain.ApiResponse{
			Code:    "911",
			Message: "SYSTEM MALFUNCTION",
		}
	}
	if vehicle == (domain.Vehicle{}) {
		return domain.ApiResponse{
			Code:    "01",
			Message: "Domain Not Found",
		}
	}

	histories, err := s.historyRepository.FindByVehicle(ctx, vehicle.ID)
	if err != nil {
		return domain.ApiResponse{
			Code:    "01",
			Message: err.Error(),
		}
	}

	var historiesData []domain.HistoryData

	for _, v := range histories {
		historiesData = append(historiesData, domain.HistoryData{
			Pic:        v.Pic,
			PlatNumber: v.PlatNumber,
			Notes:      v.Notes,
			CustomerID: v.CustomerID,
			VehicleID:  v.VehicleID,
			ComeAt:     v.CreatedAt.Format(time.RFC822Z),
		})

	}

	result := domain.VehicleHistorical{
		ID:          vehicle.ID,
		VIN:         vehicle.VIN,
		Brand:       vehicle.Brand,
		HistoryData: historiesData,
	}

	return domain.ApiResponse{
		Code:    "01",
		Message: "APPROVED",
		Data:    result,
	}
}
