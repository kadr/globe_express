package travaler_domain_models

import (
	"time"

	vo "github.com/kadr/globe_express/internal/travaler_service/domain/value_objects"

	"github.com/google/uuid"
)

type TripModel struct {
	ID            uuid.UUID
	TravelerID    uuid.UUID
	FromCountry   string
	FromCity      string
	ToCountry     string
	ToCity        string
	DepartureDate time.Time
	ArrivalDate   time.Time
	Status        string
	MaxWeightKG   float64
	MaxSizeCM3    float64
	CreatedAt     time.Time
	UpdatedAt     *time.Time
}

type TripUpdateModel struct {
	FromCountry   *string
	FromCity      *string
	ToCountry     *string
	ToCity        *string
	DepartureDate *time.Time
	ArrivalDate   *time.Time
	MaxWeightKG   *float64
	MaxSizeCM3    *float64
	Status        *string
}

func NewTrip(travelerID uuid.UUID, fromCountry, fromCity, toCountry, toCity, status string, departureDate, arrivalDate time.Time, maxWeight, maxSize float64) (TripModel, error) {
	var err error

	maxWeight, err = vo.NewWeight(maxWeight)
	if err != nil {
		return TripModel{}, err
	}
	maxSize, err = vo.NewSize(maxSize)
	if err != nil {
		return TripModel{}, err
	}
	fromCountry, err = vo.NewCountry(fromCountry)
	if err != nil {
		return TripModel{}, err
	}
	toCountry, err = vo.NewCountry(toCountry)
	if err != nil {
		return TripModel{}, err
	}
	fromCity, err = vo.NewCity(fromCity)
	if err != nil {
		return TripModel{}, err
	}
	toCity, err = vo.NewCity(toCity)
	if err != nil {
		return TripModel{}, err
	}
	status, err = vo.NewStatus(status)
	if err != nil {
		return TripModel{}, err
	}

	return TripModel{
		TravelerID:    travelerID,
		FromCountry:   fromCountry,
		FromCity:      fromCity,
		ToCountry:     toCountry,
		ToCity:        toCity,
		DepartureDate: departureDate,
		ArrivalDate:   arrivalDate,
		Status:        status,
		MaxWeightKG:   maxWeight,
		MaxSizeCM3:    maxSize,
		CreatedAt:     time.Now(),
	}, nil
}

func NewUninitializedTrip(travelerID uuid.UUID, fromCountry, fromCity, toCountry, toCity, status string, departureDate, arrivalDate time.Time, maxWeight, maxSize float64) (TripModel, error) {
	trip, err := NewTrip(travelerID, fromCountry, fromCity, toCountry, toCity, status, departureDate, arrivalDate, maxWeight, maxSize)
	if err != nil {
		return TripModel{}, err
	}
	trip.ID = uuid.UUID{}
	return trip, nil
}

func NewTripUpdateModel(fromCountry, fromCity, toCountry, toCity, status *string, departureDate, arrivalDate *time.Time, maxWeight, maxSize *float64) (TripUpdateModel, error) {
	var err error

	if maxWeight != nil {
		*maxWeight, err = vo.NewWeight(*maxWeight)
		if err != nil {
			return TripUpdateModel{}, err
		}
	}
	if maxSize != nil {
		*maxSize, err = vo.NewSize(*maxSize)
		if err != nil {
			return TripUpdateModel{}, err
		}
	}
	if fromCountry != nil {
		*fromCountry, err = vo.NewCountry(*fromCountry)
		if err != nil {
			return TripUpdateModel{}, err
		}
	}
	if toCountry != nil {
		*toCountry, err = vo.NewCountry(*toCountry)
		if err != nil {
			return TripUpdateModel{}, err
		}
	}
	if fromCity != nil {
		*fromCity, err = vo.NewCity(*fromCity)
		if err != nil {
			return TripUpdateModel{}, err
		}
	}
	if toCity != nil {
		*toCity, err = vo.NewCity(*toCity)
		if err != nil {
			return TripUpdateModel{}, err
		}
	}
	if status != nil {
		*status, err = vo.NewStatus(*status)
		if err != nil {
			return TripUpdateModel{}, err
		}
	}
	return TripUpdateModel{
		FromCountry:   fromCountry,
		FromCity:      fromCity,
		ToCountry:     toCountry,
		ToCity:        toCity,
		DepartureDate: departureDate,
		ArrivalDate:   arrivalDate,
		Status:        status,
		MaxWeightKG:   maxWeight,
		MaxSizeCM3:    maxSize,
	}, nil
}
