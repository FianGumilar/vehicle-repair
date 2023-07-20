package vehicle

import (
	"context"
	"database/sql"

	"github.com/FianGumilar/vehicle-repair/domain"
	"github.com/doug-martin/goqu/v9"
)

type repository struct {
	db *goqu.Database
}

func NewRepository(con *sql.DB) domain.VehicleRepository {
	return &repository{db: goqu.New("default", con)}
}

func (r repository) FindByVin(ctx context.Context, vin string) (vehicle domain.Vehicle, err error) {
	dataset := r.db.From("vehicles").Where(goqu.Ex{
		"vin": vin,
	}).Limit(1)

	_, err = dataset.ScanStructContext(ctx, &vehicle)
	return
}

func (r repository) FindById(ctx context.Context, id int64) (vehicle domain.Vehicle, err error) {
	dataset := r.db.From("vehicles").Where(goqu.Ex{
		"id": id,
	})

	_, err = dataset.ScanStructContext(ctx, &vehicle)
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
