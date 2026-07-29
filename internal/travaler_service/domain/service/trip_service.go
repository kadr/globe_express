package travaler_trip_service

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	domain_models "github.com/kadr/globe_express/internal/travaler_service/domain/models"
)

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
