package travaler_trip_service

import (
	"context"
	"fmt"

	domain_models "github.com/kadr/globe_express/internal/travaler_service/domain/models"
)

func (ts *TripService) GetCompleted(ctx context.Context) ([]domain_models.TripModel, error) {
	trips, err := ts.repo.GetCompleted(ctx)
	if err != nil {
		return []domain_models.TripModel{}, fmt.Errorf("trip GetActive service error: %w", err)
	}

	return trips, nil
}
