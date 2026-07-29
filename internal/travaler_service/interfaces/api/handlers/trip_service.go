package travaler_api_trip_service_interface

import (
	"context"

	"github.com/google/uuid"
	domain_models "github.com/kadr/globe_express/internal/travaler_service/domain/models"
)

type TravalerServiceIface interface {
	Create(ctx context.Context, schema domain_models.TripModel) (domain_models.TripModel, error)
	Update(ctx context.Context, tripID uuid.UUID, schema domain_models.TripUpdateModel) (domain_models.TripModel, error)
	GetActive(ctx context.Context) ([]domain_models.TripModel, error)
	GetCompleted(ctx context.Context) ([]domain_models.TripModel, error)
	Cancel(ctx context.Context, tripID uuid.UUID) error
	GetDetail(ctx context.Context, tripID uuid.UUID) (domain_models.TripModel, error)
	GetList(ctx context.Context, limit, offset int) ([]domain_models.TripModel, error)
}
