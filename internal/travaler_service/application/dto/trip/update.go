package travaler_api_trip_dto

import (
	"time"

	service "github.com/kadr/globe_express/internal/travaler_service/domain/service/trip"
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

func ToUpdateModel(schema UpdateDTO) service.TripUpdateModel {
	tripModel := service.TripUpdateModel{
		schema.FromCountry,
		schema.FromCity,
		schema.ToCountry,
		schema.ToCity,
		schema.DepartureDate,
		schema.ArrivalDate,
		schema.Status,
		schema.MaxWeightKG,
		schema.MaxSizeCM3,
	}

	return tripModel
}
