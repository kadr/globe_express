package travaler_trip_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	domain_models "github.com/kadr/globe_express/internal/travaler_service/domain/models"
)

func (ts *TripService) GetDetail(ctx context.Context, tripID uuid.UUID) (domain_models.TripModel, error) {
	trip, err := ts.repo.GetDetail(ctx, tripID)
	if err != nil {
		return domain_models.TripModel{}, fmt.Errorf("trip GetDetail service error: %w", err)
	}

	return trip, nil
}
