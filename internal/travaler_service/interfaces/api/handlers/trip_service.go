package travaler_api_trip_service_interface

import (
	"context"

	"github.com/google/uuid"
	trip_service "github.com/kadr/globe_express/internal/travaler_service/domain/service/trip"
)

type TravalerServiceIface interface {
	Create(ctx context.Context, schema trip_service.TripModel) (trip_service.TripModel, error)
	Update(ctx context.Context, tripID uuid.UUID, schema trip_service.TripUpdateModel) (trip_service.TripModel, error)
	GetActive(ctx context.Context) ([]trip_service.TripModel, error)
	GetCompleted(ctx context.Context) ([]trip_service.TripModel, error)
	Cancel(ctx context.Context, tripID uuid.UUID) error
	GetDetail(ctx context.Context, tripID uuid.UUID) (trip_service.TripModel, error)
	GetList(ctx context.Context, limit, offset int) ([]trip_service.TripModel, error)
}
