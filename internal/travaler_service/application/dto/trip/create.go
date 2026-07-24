package travaler_api_trip_dto

import (
	"time"

	"github.com/google/uuid"
	domain_models "github.com/kadr/globe_express/internal/travaler_service/domain/models"
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
}

func ToDomainModel(schema CreateDTO) (domain_models.TripModel, error) {
	// TODO: Remove travelerID
	travelerID, _ := uuid.Parse("ffaf8a49-d4a9-40ab-95c2-869b7f77f8c8")
	tripModel, err := domain_models.NewUninitializedTrip(
		travelerID,
		schema.FromCountry,
		schema.FromCity,
		schema.ToCountry,
		schema.ToCity,
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

func ToDTO(schema domain_models.TripModel) TripDTO {
	tripDTO := TripDTO{
		schema.ID.String(),
		schema.FromCountry,
		schema.FromCity,
		schema.ToCountry,
		schema.ToCity,
		schema.DepartureDate,
		schema.ArrivalDate,
		schema.MaxWeightKG,
		schema.MaxSizeCM3,
		schema.CreatedAt.Local().String(),
		nil,
	}
	if schema.UpdatedAt != nil {
		updateAt := *schema.UpdatedAt
		strUpdateAt := updateAt.Local().String()
		tripDTO.UpdatedAt = &strUpdateAt
	}

	return tripDTO
}
