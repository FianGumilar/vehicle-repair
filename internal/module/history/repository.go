package history

import (
	"context"
	"database/sql"

	"github.com/FianGumilar/vehicle-repair/domain"
)

type repository struct {
	db *sql.DB
}

func NewRepository(con *sql.DB) domain.HistoryRepository {
	return &repository{db: con}
}

func (r repository) FindByVehicle(ctx context.Context, id int64) (result []domain.HistoryDetails, err error) {

	query := `SELECT * FROM history_details WHERE vehicle_id = ? ORDER BY id ASC`
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var historyDetail domain.HistoryDetails
		var CreatedAt sql.NullTime

		err = rows.Scan(
			&historyDetail.ID,
			&historyDetail.Pic,
			&historyDetail.PlatNumber,
			&historyDetail.Notes,
			&historyDetail.CustomerID,
			&historyDetail.VehicleID,
			&CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, historyDetail)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (r repository) Insert(ctx context.Context, detail *domain.HistoryDetails) error {

	query := "INSERT INTO history_details (plate_number,pic,notes,customer_id,vehicle_id) VALUES (?, ?, ?, ?, ?)"

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
