package domain

import (
	"context"
	"time"
)

type HistoryDetails struct {
	ID         int64     `db:"id"`
	Pic        string    `db:"pic"`
	PlatNumber string    `db:"plat_number"`
	Notes      string    `db:"notes"`
	CustomerID int64     `db:"customer_id"`
	VehicleID  int64     `db:"vehicle_id"`
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
	PlatNumber string `json:"plat_number"`
	Notes      string `json:"notes"`
	CustomerID int64  `json:"customer_id"`
	VehicleID  int64  `json:"vehicle_id"`
	ComeAt     string `json:"come_at"`
}
