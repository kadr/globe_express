package travaler_api_trip_dto

import "time"

type TripDTO struct {
	ID            string    `json:"id"`
	FromCountry   string    `json:"from_country"`
	FromCity      string    `json:"from_city"`
	ToCountry     string    `json:"to_country"`
	ToCity        string    `json:"to_city"`
	DepartureDate time.Time `json:"departure_date"`
	ArrivalDate   time.Time `json:"arrival_date"`
	MaxWeightKG   float64   `json:"max_weight_kg"`
	MaxSizeCM3    float64   `json:"max_size_cm3"`
	CreatedAt     string    `json:"created_at"`
	UpdatedAt     *string   `json:"updated_at"`
}
