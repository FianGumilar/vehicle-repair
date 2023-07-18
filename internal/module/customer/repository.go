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

func (r repository) FindById(ctx context.Context, id int64) (customers domain.Customer, err error) {
	dataset := r.db.From("customers").Where(goqu.Ex{ // where the keys are string that will be used as Identifiers and values
		"id": id,
	})

	if _, err := dataset.ScanStructContext(ctx, &customers); err != nil {
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

func (r repository) FindByName(ctx context.Context, name string) (customer domain.Customer,err error) {
	dataset := r.db.From("customers").Where(goqu.Ex{ //where the keys are string that will be used as Identifiers and values
		"name": name, 
	})

	if _, err := dataset.ScanStructContext(ctx, &customer); err != nil {
		return domain.Customer{}, nil
	}

	return
}

func (r repository) FindByPhone(ctx context.Context, phone string) (customer domain.Customer,err error) {
	dataset := r.db.From("customers").Where(goqu.Ex{ // where the keys are string that will be used as Identifiers and values
		"phone": phone, 
	})

	if _, err := dataset.ScanStructContext(ctx, &customer); err != nil {
		return domain.Customer{}, nil
	}

	return
}

func (r repository) Insert(ctx context.Context, customer *domain.Customer) error {
	executor := r.db.Insert("customers").Rows(*customer).Executor()

	// Insert: retrieve table name
	// Rows: get the value stored in the variable
	// Executor: return object `goqu.Executor & execute sql`

	result, err := executor.Exec()
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	customer.ID = id

	return err
}
