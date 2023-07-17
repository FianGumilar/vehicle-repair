package domain

import (
	"context"
	"time"
)

type History struct {
	ID        int64     `db:"id"`
	NoFrame   string    `db:"no_frame"`
	Brand     string    `db:"brand"`
	UpdatedAt time.Time `db:"updated_at"`
}

type HistoryDetails struct {
	ID         int64     `db:"id"`
	Pic        string    `db:"pic"`
	PlatNumber string    `db:"plat_number"`
	Notes      string    `db:"notes"`
	CustomerID int64     `db:"customer_id"`
	HistoryID  int64     `db:"history_id"`
	CreatedAt  time.Time `db:"created_at"`
}

type HistoryRepository interface {
	FindById(ctx context.Context, id int64) (History, error)
	FindByNoFrame(ctx context.Context, no_frame string) (History, error)
	FindByHistoryDetails(ctx context.Context, id int64) ([]HistoryDetails, error)
	Insert(ctx context.Context, history *History) error
	InsertDetail(ctx context.Context, detail *HistoryDetails) error
}

type HistoryService interface {
}
