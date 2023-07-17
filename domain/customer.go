package domain

import (
	"context"
	"time"
)

type Customer struct {
	ID       int64     `db:"id"`
	Name     string    `db:"name"`
	Phone    string    `db:"phone"`
	CratedAt time.Time `db:"crated_at"`
}

type CustomerRepository interface {
	FindById(ctx context.Context, id int64) (Customer, error)
	FindByIds(ctx context.Context, id []int64) ([]Customer, error)
	FindByName(ctx context.Context, name string) (Customer, error)
	FindByPhone(ctx context.Context, phone string) (Customer, error)
	Insert(ctx context.Context, customer *Customer) error
}

type CustomerService interface {
}
