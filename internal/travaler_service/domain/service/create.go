package travaler_trip_service

import (
	"context"
	"fmt"

	domain_models "github.com/kadr/globe_express/internal/travaler_service/domain/models"
)

func (ts *TripService) Create(ctx context.Context, schema domain_models.TripModel) (domain_models.TripModel, error) {
	trip, err := ts.repo.Create(ctx, schema)
	if err != nil {
		return domain_models.TripModel{}, fmt.Errorf("trip service error: %w", err)
	}

	return trip, nil
}
