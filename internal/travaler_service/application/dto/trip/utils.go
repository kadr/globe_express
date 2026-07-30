package travaler_api_trip_dto

import (
	trip_service "github.com/kadr/globe_express/internal/travaler_service/domain/service/trip"
)

func ToDTO(schema trip_service.TripModel) TripDTO {
	tripDTO := TripDTO{
		schema.ID.String(),
		schema.FromCountry,
		schema.FromCity,
		schema.ToCountry,
		schema.ToCity,
		schema.DepartureDate,
		schema.ArrivalDate,
		*schema.Status,
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

func ToDTOList(schema []trip_service.TripModel) []TripDTO {
	var results []TripDTO
	for _, trip := range schema {
		results = append(results, ToDTO(trip))
	}

	return results
}
