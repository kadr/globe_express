package travaler_trip_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (ts *TripService) GetDetail(ctx context.Context, tripID uuid.UUID) (TripModel, error) {
	trip, err := ts.repo.GetDetail(ctx, tripID)
	if err != nil {
		return TripModel{}, fmt.Errorf("trip GetDetail service error: %w", err)
	}

	return FromDomainModel(trip), nil
}
