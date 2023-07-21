package vehicle

import (
	"context"
	"database/sql"

	"github.com/FianGumilar/vehicle-repair/domain"
)

type repository struct {
	db *sql.DB
}

func NewRepository(con *sql.DB) domain.VehicleRepository {
	return &repository{db: con}
}

func (r repository) FindByVin(ctx context.Context, vin string) (domain.Vehicle, error) {
	query := `SELECT * FROM vehicles WHERE vin = ? LIMIT 1`

	var vehicle domain.Vehicle

	err := r.db.QueryRowContext(ctx, query, vin).Scan(
		&vehicle.ID,
		&vehicle.VIN,
		&vehicle.Brand,
		&vehicle.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return vehicle, nil
	}

	return vehicle, err
}

func (r repository) FindById(ctx context.Context, id int64) (vehicle domain.Vehicle, err error) {
	query := `SELECT * FROM vehicles WHERE id = ? LIMIT 1`
	err = r.db.QueryRowContext(ctx, query, id).Scan(
		&vehicle.ID,
		&vehicle.VIN,
		&vehicle.Brand,
		&vehicle.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return vehicle, nil
	}
	return
}

func (r repository) Insert(ctx context.Context, vehicle *domain.Vehicle) error {
	// Query SQL untuk melakukan insert data ke tabel "vehicles"
	query := `INSERT INTO vehicles (vin, brand, updated_at) VALUES (?, ?, ?)`

	// Eksekusi query SQL menggunakan prepared statement
	result, err := r.db.ExecContext(ctx, query, vehicle.VIN, vehicle.Brand, vehicle.UpdatedAt)
	if err != nil {
		return err
	}

	// Ambil ID dari data yang baru saja diinsert
	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Setel ID yang baru saja diinsert ke struktur kendaraan
	vehicle.ID = lastInsertedID
	return nil
}
