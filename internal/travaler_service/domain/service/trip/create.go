package travaler_trip_service

import (
	"context"
	"fmt"

	domain_models "github.com/kadr/globe_express/internal/travaler_service/domain/models"
)

func (ts *TripService) Create(ctx context.Context, schema TripModel) (TripModel, error) {
	domainSchema, err := ToDomainModel(schema)
	if err != nil {
		return TripModel{}, fmt.Errorf("trip service error: %w", err)
	}
	trip, err := ts.repo.Create(ctx, domainSchema)
	if err != nil {
		return TripModel{}, fmt.Errorf("trip service error: %w", err)
	}

	return FromDomainModel(trip), nil
}

func ToDomainModel(schema TripModel) (domain_models.TripModel, error) {
	newStatus := ""
	if schema.Status != nil {
		newStatus = *schema.Status
	}
	tripModel, err := domain_models.NewUninitializedTrip(
		schema.TravelerID,
		schema.FromCountry,
		schema.FromCity,
		schema.ToCountry,
		schema.ToCity,
		newStatus,
		schema.DepartureDate,
		schema.ArrivalDate,
		schema.MaxWeightKG,
		schema.MaxSizeCM3,
	)
	if err != nil {
		return domain_models.TripModel{}, err
	}

	return tripModel, nil
}
