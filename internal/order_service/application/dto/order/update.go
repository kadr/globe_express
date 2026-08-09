package order_api_order_dto

import (
	"time"

	"github.com/google/uuid"
	services "github.com/kadr/globe_express/internal/order_service/interfaces/api/handlers"
)

type UpdateDTO struct {
	TravelerID       *uuid.UUID `json:"traveler_id"`
	Status           *string    `json:"status"`
	PickupCity       *string    `json:"pickup_city"`
	PickupCountry    *string    `json:"pickup_country"`
	DeliveryCity     *string    `json:"delivery_city"`
	DeliveryCountry  *string    `json:"delivery_country"`
	RewardAmount     *float64   `json:"reward_amount"`
	RewardCurrency   *string    `json:"reward_currency"`
	DeliveryDeadline *time.Time `json:"delivery_deadline"`
	DeliveryDate     *time.Time `json:"delivery_date"`
}

func ToUpdateModel(schema UpdateDTO) services.OrderUpdateModel {
	orderModel := services.OrderUpdateModel{
		schema.TravelerID,
		schema.Status,
		schema.PickupCity,
		schema.PickupCountry,
		schema.DeliveryCity,
		schema.DeliveryCountry,
		schema.RewardAmount,
		schema.RewardCurrency,
		schema.DeliveryDeadline,
		schema.DeliveryDate,
	}

	return orderModel
}
