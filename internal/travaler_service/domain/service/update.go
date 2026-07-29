package travaler_trip_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	domain_models "github.com/kadr/globe_express/internal/travaler_service/domain/models"
)

func (ts *TripService) Update(ctx context.Context, tripID uuid.UUID, schema domain_models.TripUpdateModel) (domain_models.TripModel, error) {
	trip, err := ts.repo.Update(ctx, tripID, schema)
	if err != nil {
		return domain_models.TripModel{}, fmt.Errorf("trip update service error: %w", err)
	}

	return trip, nil
}
