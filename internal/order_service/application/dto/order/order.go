package order_api_order_dto

import (
	"time"

	product_dto "github.com/kadr/globe_express/internal/order_service/application/dto/product"
)

type OrderDTO struct {
	ID               string                    `json:"id"`
	CustomerID       string                    `json:"customer_id"`
	TravelerID       *string                   `json:"traveler_id"`
	Status           string                    `json:"status"`
	PickupCity       string                    `json:"pickup_city"`
	PickupCountry    string                    `json:"pickup_country"`
	DeliveryCity     string                    `json:"delivery_city"`
	DeliveryCountry  string                    `json:"delivery_country"`
	RewardAmount     float64                   `json:"reward_amount"`
	RewardCurrency   string                    `json:"reward_currency"`
	DeliveryDeadline time.Time                 `json:"delivery_deadline"`
	DeliveryDate     time.Time                 `json:"delivery_date"`
	CreatedAt        string                    `json:"created_at"`
	UpdatedAt        *string                   `json:"updated_at"`
	Products         *[]product_dto.ProductDTO `json:"products"`
}
