package travaler_api_trip_dto

import (
	"time"

	"github.com/google/uuid"
	service "github.com/kadr/globe_express/internal/travaler_service/domain/service/trip"
)

type CreateDTO struct {
	FromCountry   string    `json:"from_country"`
	FromCity      string    `json:"from_city"`
	ToCountry     string    `json:"to_country"`
	ToCity        string    `json:"to_city"`
	DepartureDate time.Time `json:"departure_date"`
	ArrivalDate   time.Time `json:"arrival_date"`
	MaxWeightKG   float64   `json:"max_weight_kg"`
	MaxSizeCM3    float64   `json:"max_size_cm3"`
	Status        *string   `json:"status"`
}

func ToCreateModel(schema CreateDTO, userID uuid.UUID) service.TripModel {
	tripModel := service.TripModel{
		nil,
		userID,
		schema.FromCountry,
		schema.FromCity,
		schema.ToCountry,
		schema.ToCity,
		schema.DepartureDate,
		schema.ArrivalDate,
		schema.Status,
		schema.MaxWeightKG,
		schema.MaxSizeCM3,
		nil,
		nil,
	}

	return tripModel
}
