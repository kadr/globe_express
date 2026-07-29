package travaler_trip_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	domain_models "github.com/kadr/globe_express/internal/travaler_service/domain/models"
	vo "github.com/kadr/globe_express/internal/travaler_service/domain/value_objects"
	api_errors "github.com/kadr/globe_express/pkg/errors"
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
	Status        string     `db:"status"`
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
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(tr.timeout))
	defer cancel()
	var result TripResult
	stmt := `
	INSERT INTO trips(traveler_id,from_country,from_city,to_country,to_city,departure_date,arrival_date,status,max_weight_kg,max_size_cm3)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) 
	RETURNING id,traveler_id,from_country,from_city,to_country,to_city,departure_date,arrival_date,status,max_weight_kg,max_size_cm3,created_at,updated_at
	`
	row := tr.db.QueryRowxContext(ctx, stmt,
		schema.TravelerID,
		schema.FromCountry,
		schema.FromCity,
		schema.ToCountry,
		schema.ToCity,
		schema.DepartureDate,
		schema.ArrivalDate,
		schema.Status,
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

func (tr *TripRepository) Update(ctx context.Context, tripID uuid.UUID, schema domain_models.TripUpdateModel) (domain_models.TripModel, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(tr.timeout))
	defer cancel()
	var setParts []string
	args := map[string]interface{}{"id": tripID}
	if schema.FromCountry != nil {
		setParts = append(setParts, "from_country=:from_country")
		args["from_country"] = schema.FromCountry
	}
	if schema.FromCity != nil {
		setParts = append(setParts, "from_city=:from_city")
		args["from_city"] = schema.FromCity
	}
	if schema.ToCountry != nil {
		setParts = append(setParts, "to_country=:to_country")
		args["to_country"] = schema.ToCountry
	}
	if schema.ToCity != nil {
		setParts = append(setParts, "to_city=:to_city")
		args["to_city"] = schema.ToCity
	}
	if schema.DepartureDate != nil {
		setParts = append(setParts, "departure_date=:departure_date")
		args["departure_date"] = schema.DepartureDate
	}
	if schema.ArrivalDate != nil {
		setParts = append(setParts, "arrival_date=:arrival_date")
		args["arrival_date"] = schema.ArrivalDate
	}
	if schema.Status != nil {
		setParts = append(setParts, "status=:status")
		args["status"] = schema.Status
	}
	if schema.MaxWeightKG != nil {
		setParts = append(setParts, "max_weight_kg=:max_weight_kg")
		args["max_weight_kg"] = schema.MaxWeightKG
	}
	if schema.MaxSizeCM3 != nil {
		setParts = append(setParts, "max_size_cm3=:max_size_cm3")
		args["max_size_cm3"] = schema.MaxSizeCM3
	}

	stmt := fmt.Sprintf(
		`
		UPDATE trips SET %s, updated_at=NOW() WHERE id=:id 
		RETURNING id,traveler_id,from_country,from_city,to_country,to_city,departure_date,arrival_date,status,max_weight_kg,max_size_cm3,created_at,updated_at`,
		strings.Join(setParts, ","),
	)
	rows, err := tr.db.NamedQueryContext(ctx, stmt, args)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			tr.logger.Warn(err.Error())
			return domain_models.TripModel{}, fmt.Errorf("trip repository, update record no found: %w", api_errors.ErrorNotFound)
		}
		tr.logger.Error(err.Error())
		return domain_models.TripModel{}, fmt.Errorf("trip repository, update err: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		var result TripResult
		err := rows.StructScan(&result)
		if err != nil {
			tr.logger.Error(err.Error())
			return domain_models.TripModel{}, fmt.Errorf("trip repository, update err: %w", err)
		}

		return toDomainModel(result), nil
	}
	return domain_models.TripModel{}, fmt.Errorf("trip repository, update err: can't update record")
}

func (tr *TripRepository) GetActive(ctx context.Context) ([]domain_models.TripModel, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(tr.timeout))
	defer cancel()
	query := `
	SELECT id,traveler_id,from_country,from_city,to_country,to_city,departure_date,arrival_date,status,max_weight_kg,max_size_cm3,created_at,updated_at
	FROM trips WHERE status != $1 ORDER BY created_at DESC`

	var results []TripResult
	err := tr.db.SelectContext(ctx, &results, query, vo.Completed)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			tr.logger.Warn(err.Error())
			return []domain_models.TripModel{}, nil
		}
		tr.logger.Error(err.Error())
		return []domain_models.TripModel{}, fmt.Errorf("trip repository, GetActive err: %w", err)
	}

	return toDomainModelList(results), nil
}

func (tr *TripRepository) GetCompleted(ctx context.Context) ([]domain_models.TripModel, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(tr.timeout))
	defer cancel()
	query := `
	SELECT id,traveler_id,from_country,from_city,to_country,to_city,departure_date,arrival_date,status,max_weight_kg,max_size_cm3,created_at,updated_at
	FROM trips WHERE status = $1 ORDER BY created_at DESC`

	var results []TripResult
	err := tr.db.SelectContext(ctx, &results, query, vo.Completed)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			tr.logger.Warn(err.Error())
			return []domain_models.TripModel{}, nil
		}
		tr.logger.Error(err.Error())
		return []domain_models.TripModel{}, fmt.Errorf("trip repository, GetActive err: %w", err)
	}

	return toDomainModelList(results), nil
}

func (tr *TripRepository) Cancel(ctx context.Context, tripID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(tr.timeout))
	defer cancel()
	stmt := "DELETE FROM trips WHERE id = $1"

	res, err := tr.db.ExecContext(ctx, stmt, tripID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			tr.logger.Warn(err.Error())
			return fmt.Errorf("trip repository, record no found: %w", api_errors.ErrorNotFound)
		}
		tr.logger.Error(err.Error())
		return fmt.Errorf("trip repository, Cancel err: %w", err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("trip repository, Cancel err: %w", err)
	}
	if count == 0 {
		tr.logger.Warn("no rows found to delete")
	}

	return nil
}

func (tr *TripRepository) GetDetail(ctx context.Context, tripID uuid.UUID) (domain_models.TripModel, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(tr.timeout))
	defer cancel()
	query := `
	SELECT id,traveler_id,from_country,from_city,to_country,to_city,departure_date,arrival_date,status,max_weight_kg,max_size_cm3,created_at,updated_at
	FROM trips WHERE id = $1 ORDER BY created_at DESC`

	var result TripResult
	err := tr.db.GetContext(ctx, &result, query, tripID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			tr.logger.Warn(err.Error())
			return domain_models.TripModel{}, fmt.Errorf("trip repository, record no found: %w", api_errors.ErrorNotFound)
		}
		tr.logger.Error(err.Error())
		return domain_models.TripModel{}, fmt.Errorf("trip repository, GetDetail err: %w", err)
	}

	return toDomainModel(result), nil
}

func (tr *TripRepository) GetList(ctx context.Context, limit, offset int) ([]domain_models.TripModel, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(tr.timeout))
	defer cancel()
	query := `
	SELECT id,traveler_id,from_country,from_city,to_country,to_city,departure_date,arrival_date,status,max_weight_kg,max_size_cm3,created_at,updated_at
	FROM trips ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	var results []TripResult
	err := tr.db.SelectContext(ctx, &results, query, limit, offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			tr.logger.Warn(err.Error())
			return []domain_models.TripModel{}, nil
		}
		tr.logger.Error(err.Error())
		return []domain_models.TripModel{}, fmt.Errorf("trip repository, GetList err: %w", err)
	}

	return toDomainModelList(results), nil
}

func toDomainModel(result TripResult) domain_models.TripModel {
	return domain_models.TripModel(result)
}

func toDomainModelList(trips []TripResult) []domain_models.TripModel {
	var results []domain_models.TripModel
	for _, res := range trips {
		results = append(results, domain_models.TripModel(res))
	}
	return results
}
