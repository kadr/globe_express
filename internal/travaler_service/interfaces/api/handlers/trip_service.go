package travaler_api_trip_service_interface

import (
	"context"

	domain_models "github.com/kadr/globe_express/internal/travaler_service/domain/models"
)

type TravalerServiceIface interface {
	Create(ctx context.Context, schema domain_models.TripModel) (domain_models.TripModel, error)
}
