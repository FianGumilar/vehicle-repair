package domain

import (
	"context"
	"time"
)

type Customer struct {
	ID       int64     `db:"id"`
	Name     string    `db:"name" validate:"required"`
	Phone    string    `db:"phone" validate:"required"`
	CretedAt time.Time `db:"crated_at"`
}

type CustomerData struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type CustomerRepository interface {
	FindAll(ctx context.Context) ([]Customer, error)
	FindById(ctx context.Context, id int64) (Customer, error)
	FindByIds(ctx context.Context, id []int64) ([]Customer, error)
	Insert(ctx context.Context, customer *Customer) error
}

type CustomerService interface {
	All(ctx context.Context) ApiResponse
	Save(ctx context.Context, customer CustomerData) ApiResponse
}
