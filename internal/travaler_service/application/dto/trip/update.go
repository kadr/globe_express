package travaler_api_trip_dto

import (
	"time"

	domain_models "github.com/kadr/globe_express/internal/travaler_service/domain/models"
)

type UpdateDTO struct {
	FromCountry   *string    `json:"from_country"`
	FromCity      *string    `json:"from_city"`
	ToCountry     *string    `json:"to_country"`
	ToCity        *string    `json:"to_city"`
	DepartureDate *time.Time `json:"departure_date"`
	ArrivalDate   *time.Time `json:"arrival_date"`
	Status        *string    `json:"status"`
	MaxWeightKG   *float64   `json:"max_weight_kg"`
	MaxSizeCM3    *float64   `json:"max_size_cm3"`
}

func ToUpdateDomainModel(schema UpdateDTO) (domain_models.TripUpdateModel, error) {
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
