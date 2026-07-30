package travaler_trip_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	vo "github.com/kadr/globe_express/internal/travaler_service/domain/value_objects"
	api_errors "github.com/kadr/globe_express/pkg/errors"
)

func (ts *TripService) Cancel(ctx context.Context, tripID uuid.UUID) error {
	if trip, err := ts.repo.GetDetail(ctx, tripID); err == nil {
		if trip.Status != vo.Planned {
			return api_errors.ErrorCanNotCancel
		}
	}
	err := ts.repo.Cancel(ctx, tripID)
	if err != nil {
		return fmt.Errorf("trip Cancel service error: %w", err)
	}

	return nil
}
