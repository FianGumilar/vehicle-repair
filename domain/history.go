package domain

import (
	"context"
	"time"
)

type HistoryDetails struct {
	ID         int64     `db:"id"`
	Pic        string    `db:"pic" form:"pic" validate:"required"`
	PlatNumber string    `db:"plate_number" form:"plate_number" validate:"required"`
	Notes      string    `db:"notes" form:"notes" validate:"required"`
	CustomerID int64     `db:"customer_id" form:"customer_id" validate:"required"`
	VehicleID  int64     `db:"vehicle_id" form:"vehicle_id" validate:"required"`
	CreatedAt  time.Time `db:"created_at"`
}

type HistoryRepository interface {
	FindByVehicle(ctx context.Context, id int64) ([]HistoryDetails, error)
	Insert(ctx context.Context, detail *HistoryDetails) error
}

type HistoryService interface {
}

type HistoryData struct {
	ID         int64  `json:"id"`
	Pic        string `json:"pic"`
	PlatNumber string `json:"plate_number"`
	Notes      string `json:"notes"`
	CustomerID int64  `json:"customer_id"`
	VehicleID  int64  `json:"vehicle_id"`
	ComeAt     string `json:"come_at"`
}
