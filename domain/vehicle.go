package domain

import (
	"context"
	"time"
)

type Vehicle struct {
	ID        int64     `db:"id"`
	VIN       string    `db:"vin"`
	Brand     string    `db:"brand"`
	UpdatedAt time.Time `db:"updated_at"`
}

type VehicleRepository interface {
	FindById(ctx context.Context, id int64) (Vehicle, error)
	FindByVin(ctx context.Context, vin string) (Vehicle, error)
	Insert(ctx context.Context, vehicle *Vehicle) error
}

type VehicleService interface {
	FindHistorical(ctx context.Context, vin string) ApiResponse
}

type VehicleHistorical struct {
	ID          int64         `json:"id"`
	VIN         string        `json:"vin"`
	Brand       string        `json:"brand"`
	HistoryData []HistoryData `json:"history_data"`
}
