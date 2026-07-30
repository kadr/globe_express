package travaler_trip_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	domain_models "github.com/kadr/globe_express/internal/travaler_service/domain/models"
)

func (ts *TripService) Update(ctx context.Context, tripID uuid.UUID, schema TripUpdateModel) (TripModel, error) {
	newSchema, err := ToUpdateDomainModel(schema)
	if err != nil {
		return TripModel{}, fmt.Errorf("trip update service error: %w", err)
	}
	trip, err := ts.repo.Update(ctx, tripID, newSchema)
	if err != nil {
		return TripModel{}, fmt.Errorf("trip update service error: %w", err)
	}

	return FromDomainModel(trip), nil
}

func ToUpdateDomainModel(schema TripUpdateModel) (domain_models.TripUpdateModel, error) {
	tripModel, err := domain_models.NewTripUpdateModel(
		schema.FromCountry,
		schema.FromCity,
		schema.ToCountry,
		schema.ToCity,
		schema.Status,
		schema.DepartureDate,
		schema.ArrivalDate,
		schema.MaxWeightKG,
		schema.MaxSizeCM3,
	)
	if err != nil {
		return domain_models.TripUpdateModel{}, err
	}

	return tripModel, nil
}
