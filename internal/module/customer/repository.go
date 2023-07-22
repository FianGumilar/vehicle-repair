package customer

import (
	"context"
	"database/sql"

	"github.com/FianGumilar/vehicle-repair/domain"
)

type repository struct {
	db *sql.DB
}

func NewRepository(con *sql.DB) domain.CustomerRepository {
	return &repository{db: con}
}

func (r repository) FindAll(ctx context.Context) (customers []domain.Customer, err error) {
	query := `SELECT * FROM customers ORDER BY name ASC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer domain.Customer

		err := rows.Scan(&customer.ID, &customer.Name, &customer.Phone, &customer.CretedAt)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return customers, nil
}

func (r repository) FindById(ctx context.Context, id int64) (customers domain.Customer, err error) {
	query := `SELECT * FROM customers WHERE id = ? LIMIT 1`

	err = r.db.QueryRowContext(ctx, query, id).Scan(
		&customers.ID,
		&customers.Name,
		&customers.Phone,
		&customers.CretedAt,
	)
	if err == sql.ErrNoRows {
		return customers, nil
	}
	return
}

func (r repository) FindByIds(ctx context.Context, ids []int64) (customers []domain.Customer, err error) {
	query := `SELECT * FROM customers WHERE ids = ? LIMIT 1`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer domain.Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.Phone, &customer.CretedAt)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return customers, nil
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
	err = stmt.QueryRowContext(ctx, customer.Name, customer.Phone, customer.CretedAt).Scan(&id)
	if err != nil {
		return err
	}

	customer.ID = id
	return nil
}
