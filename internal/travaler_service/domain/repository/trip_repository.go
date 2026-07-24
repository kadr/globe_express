package travaler_trip_repository

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	domain_models "github.com/kadr/globe_express/internal/travaler_service/domain/models"
	_ "github.com/lib/pq"
)

type TripResult struct {
	ID            uuid.UUID  `db:"id"`
	TravelerID    uuid.UUID  `db:"traveler_id"`
	FromCountry   string     `db:"from_country"`
	FromCity      string     `db:"from_city"`
	ToCountry     string     `db:"to_country"`
	ToCity        string     `db:"to_city"`
	DepartureDate time.Time  `db:"departure_date"`
	ArrivalDate   time.Time  `db:"arrival_date"`
	MaxWeightKG   float64    `db:"max_weight_kg"`
	MaxSizeCM3    float64    `db:"max_size_cm3"`
	CreatedAt     time.Time  `db:"created_at"`
	UpdatedAt     *time.Time `db:"updated_at"`
}

type TripRepository struct {
	db      *sqlx.DB
	timeout int
	logger  *slog.Logger
}

func NewRepository(db *sqlx.DB, timeout int, logger *slog.Logger) *TripRepository {
	return &TripRepository{db: db, timeout: timeout, logger: logger}
}

func (tr *TripRepository) Create(ctx context.Context, schema domain_models.TripModel) (domain_models.TripModel, error) {
	tr.logger.Debug("timeout", tr.timeout)
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(tr.timeout))
	defer cancel()
	var result TripResult
	stmt := `
	INSERT INTO trips(traveler_id,from_country,from_city,to_country,to_city,departure_date,arrival_date,max_weight_kg,max_size_cm3)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) 
	RETURNING id,traveler_id,from_country,from_city,to_country,to_city,departure_date,arrival_date,max_weight_kg,max_size_cm3,created_at,updated_at
	`
	row := tr.db.QueryRowxContext(ctx, stmt,
		schema.TravelerID,
		schema.FromCountry,
		schema.FromCity,
		schema.ToCountry,
		schema.ToCity,
		schema.DepartureDate,
		schema.ArrivalDate,
		schema.MaxWeightKG,
		schema.MaxSizeCM3,
	)
	err := row.StructScan(&result)
	if err != nil {
		tr.logger.Error(err.Error())
		return domain_models.TripModel{}, fmt.Errorf("trip repository, create err: %w", err)
	}

	return toDomainModel(result), nil
}

func toDomainModel(result TripResult) domain_models.TripModel {
	return domain_models.TripModel(result)
}
