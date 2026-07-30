package travaler_trip_service

import (
	"context"
	"fmt"
)

func (ts *TripService) GetList(ctx context.Context, limit, offset int) ([]TripModel, error) {
	trips, err := ts.repo.GetList(ctx, limit, offset)
	if err != nil {
		return []TripModel{}, fmt.Errorf("trip GetList service error: %w", err)
	}

	return FromDomainModelList(trips), nil
}
