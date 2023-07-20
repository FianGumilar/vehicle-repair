package history

import (
	"context"
	"database/sql"
	"time"

	"github.com/FianGumilar/vehicle-repair/domain"
	"github.com/doug-martin/goqu/v9"
)

type repository struct {
	db *goqu.Database
}

func NewRepository(con *sql.DB) domain.HistoryRepository {
	return &repository{db: goqu.New("default", con)}
}

func (r repository) FindByVehicle(ctx context.Context, id int64) (result []domain.HistoryDetails, err error) {
	dataset := r.db.From("history_details").Where(goqu.Ex{
		"vehicle_id": id,
	}).Order(goqu.I("id").Asc())

	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

func (r repository) Insert(ctx context.Context, detail *domain.HistoryDetails) error {
	detail.CreatedAt = time.Now().UTC()

	query := "INSERT INTO history_details (pic, plat_number, notes, customer_id, vehicle_id) VALUES (?, ?, ?, ?, ?)"

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, detail.Pic, detail.PlatNumber, detail.Notes, detail.CustomerID, detail.VehicleID)
	if err != nil {
		return err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	detail.ID = insertedID

	return nil
}
