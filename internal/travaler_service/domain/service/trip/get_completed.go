package travaler_trip_service

import (
	"context"
	"fmt"
)

func (ts *TripService) GetCompleted(ctx context.Context) ([]TripModel, error) {
	trips, err := ts.repo.GetCompleted(ctx)
	if err != nil {
		return []TripModel{}, fmt.Errorf("trip GetActive service error: %w", err)
	}

	return FromDomainModelList(trips), nil
}
