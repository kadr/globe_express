package travaler_trip_service

import (
	"context"
	"fmt"

	domain_models "github.com/kadr/globe_express/internal/travaler_service/domain/models"
)

func (ts *TripService) GetList(ctx context.Context, limit, offset int) ([]domain_models.TripModel, error) {
	trips, err := ts.repo.GetList(ctx, limit, offset)
	if err != nil {
		return []domain_models.TripModel{}, fmt.Errorf("trip GetList service error: %w", err)
	}

	return trips, nil
}
