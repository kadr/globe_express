package travaler_trip_service

import (
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"
	domain_models "github.com/kadr/globe_express/internal/travaler_service/domain/models"
)

type TripModel struct {
	ID            *uuid.UUID
	TravelerID    uuid.UUID
	FromCountry   string
	FromCity      string
	ToCountry     string
	ToCity        string
	DepartureDate time.Time
	ArrivalDate   time.Time
	Status        *string
	MaxWeightKG   float64
	MaxSizeCM3    float64
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}

type TripUpdateModel struct {
	FromCountry   *string
	FromCity      *string
	ToCountry     *string
	ToCity        *string
	DepartureDate *time.Time
	ArrivalDate   *time.Time
	Status        *string
	MaxWeightKG   *float64
	MaxSizeCM3    *float64
}

type TripRepositoryIface interface {
	Create(ctx context.Context, schema domain_models.TripModel) (domain_models.TripModel, error)
	Update(ctx context.Context, tripID uuid.UUID, schema domain_models.TripUpdateModel) (domain_models.TripModel, error)
	GetActive(ctx context.Context) ([]domain_models.TripModel, error)
	GetCompleted(ctx context.Context) ([]domain_models.TripModel, error)
	Cancel(ctx context.Context, tripID uuid.UUID) error
	GetDetail(ctx context.Context, tripID uuid.UUID) (domain_models.TripModel, error)
	GetList(ctx context.Context, limit, offset int) ([]domain_models.TripModel, error)
}

type TripService struct {
	repo   TripRepositoryIface
	logger *slog.Logger
}

func NewService(repo TripRepositoryIface, logger *slog.Logger) *TripService {
	return &TripService{repo: repo, logger: logger}
}

func FromDomainModel(schema domain_models.TripModel) TripModel {
	tripModel := TripModel{
		&schema.ID,
		schema.TravelerID,
		schema.FromCountry,
		schema.FromCity,
		schema.ToCountry,
		schema.ToCity,
		schema.DepartureDate,
		schema.ArrivalDate,
		&schema.Status,
		schema.MaxWeightKG,
		schema.MaxSizeCM3,
		&schema.CreatedAt,
		schema.UpdatedAt,
	}

	return tripModel
}

func FromDomainModelList(result []domain_models.TripModel) []TripModel {
	var results []TripModel
	for _, trip := range result {
		results = append(results, TripModel{
			&trip.ID,
			trip.TravelerID,
			trip.FromCountry,
			trip.FromCity,
			trip.ToCountry,
			trip.ToCity,
			trip.DepartureDate,
			trip.ArrivalDate,
			&trip.Status,
			trip.MaxWeightKG,
			trip.MaxSizeCM3,
			&trip.CreatedAt,
			trip.UpdatedAt,
		})
	}

	return results
}
