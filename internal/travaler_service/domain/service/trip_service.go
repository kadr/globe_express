package travaler_trip_service

import (
	"context"
	"log/slog"

	domain_models "github.com/kadr/globe_express/internal/travaler_service/domain/models"
)

type TripRepositoryIface interface {
	Create(ctx context.Context, schema domain_models.TripModel) (domain_models.TripModel, error)
}

type TripService struct {
	repo   TripRepositoryIface
	logger *slog.Logger
}

func NewService(repo TripRepositoryIface, logger *slog.Logger) *TripService {
	return &TripService{repo: repo, logger: logger}
}
