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

	executor := r.db.Insert("history_details").Rows(goqu.Record{
		"pic":         detail.Pic,
		"plat_number": detail.PlatNumber,
		"notes":       detail.Notes,
		"customer_id": detail.CustomerID,
		"vehicle_id":  detail.VehicleID,
		"created_id":  detail.CreatedAt,
	}).Returning("id").Executor()

	_, err := executor.ScanStructContext(ctx, detail)
	return err
}
