package customer

import (
	"context"
	"database/sql"

	"github.com/FianGumilar/vehicle-repair/domain"
	"github.com/doug-martin/goqu/v9"
)

type repository struct {
	db *goqu.Database
}

func NewRepository(con *sql.DB) domain.CustomerRepository {
	return &repository{db: goqu.New("default", con)}
}

func (r repository) FindAll(ctx context.Context) (customers []domain.Customer, err error) {
	dataset := r.db.From("customers").Order(goqu.I("name").Asc())

	if errScan := dataset.ScanStructsContext(ctx, &customers); err != nil {
		return nil, errScan
	}
	return

}

func (r repository) FindById(ctx context.Context, id int64) (customers domain.Customer, err error) {
	dataset := r.db.From("customers").Where(goqu.Ex{ // where the keys are string that will be used as Identifiers and values
		"id": id,
	})

	if _, errScan := dataset.ScanStructContext(ctx, &customers); errScan != nil {
		return domain.Customer{}, nil
	}

	return
}

func (r repository) FindByIds(ctx context.Context, ids []int64) (customer []domain.Customer, err error) {
	dataset := r.db.From("customers").Where(goqu.Ex{ // where the keys are string that will be used as Identifiers and values
		"id": ids,
	})

	if err := dataset.ScanStructsContext(ctx, &customer); err != nil {
		return nil, err
	}

	return
}

func (r repository) FindByName(ctx context.Context, name string) (customer domain.Customer, err error) {
	dataset := r.db.From("customers").Where(goqu.Ex{ //where the keys are string that will be used as Identifiers and values
		"name": name,
	})

	if _, err := dataset.ScanStructContext(ctx, &customer); err != nil {
		return domain.Customer{}, nil
	}

	return
}

func (r repository) FindByPhone(ctx context.Context, phone string) (customer domain.Customer, err error) {
	dataset := r.db.From("customers").Where(goqu.Ex{ // where the keys are string that will be used as Identifiers and values
		"phone": phone,
	})

	if _, err := dataset.ScanStructContext(ctx, &customer); err != nil {
		return domain.Customer{}, nil
	}

	return
}

func (r repository) Insert(ctx context.Context, customer *domain.Customer) error {
	query := `
		INSERT INTO customers (name, phone, created_at)
		VALUES (?, ?, ?)
		RETURNING id
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRowContext(ctx, customer.Name, customer.Phone, customer.CratedAt).Scan(&id)
	if err != nil {
		return err
	}

	customer.ID = id
	return nil
}
