package travaler_api_trip_dto

import (
	domain_models "github.com/kadr/globe_express/internal/travaler_service/domain/models"
)

func ToDTO(schema domain_models.TripModel) TripDTO {
	tripDTO := TripDTO{
		schema.ID.String(),
		schema.FromCountry,
		schema.FromCity,
		schema.ToCountry,
		schema.ToCity,
		schema.DepartureDate,
		schema.ArrivalDate,
		schema.Status,
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

func ToDTOList(schema []domain_models.TripModel) []TripDTO {
	var results []TripDTO
	for _, trip := range schema {
		dto := TripDTO{
			trip.ID.String(),
			trip.FromCountry,
			trip.FromCity,
			trip.ToCountry,
			trip.ToCity,
			trip.DepartureDate,
			trip.ArrivalDate,
			trip.Status,
			trip.MaxWeightKG,
			trip.MaxSizeCM3,
			trip.CreatedAt.Local().String(),
			nil,
		}
		if trip.UpdatedAt != nil {
			updateAt := *trip.UpdatedAt
			strUpdateAt := updateAt.Local().String()
			dto.UpdatedAt = &strUpdateAt
		}
		results = append(results, dto)
	}

	return results
}
