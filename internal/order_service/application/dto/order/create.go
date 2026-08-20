package order_api_order_dto

import (
	"time"

	"github.com/google/uuid"
	product_dto "github.com/kadr/globe_express/internal/order_service/application/dto/product"
	product_service "github.com/kadr/globe_express/internal/order_service/domain/service/product"
	services "github.com/kadr/globe_express/internal/order_service/interfaces/api/handlers"
)

type CreateDTO struct {
	CustomerID       uuid.UUID                `json:"customer_id"`
	TravelerID       *uuid.UUID               `json:"traveler_id"`
	Status           *string                  `json:"status"`
	PickupCity       string                   `json:"pickup_city"`
	PickupCountry    string                   `json:"pickup_country"`
	DeliveryCity     string                   `json:"delivery_city"`
	DeliveryCountry  string                   `json:"delivery_country"`
	RewardAmount     float64                  `json:"reward_amount"`
	RewardCurrency   string                   `json:"reward_currency"`
	DeliveryDeadline time.Time                `json:"delivery_deadline"`
	DeliveryDate     time.Time                `json:"delivery_date"`
	Products         *[]product_dto.CreateDTO `json:"products"`
}

func ToCreateModel(schema CreateDTO) services.OrderModel {
	orderModel := services.OrderModel{
		nil,
		schema.CustomerID,
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
		nil,
		nil,
		nil,
	}
	if schema.Products != nil {
		products := make([]product_service.ProductModel, 0, len(*schema.Products))
		orderModel.Products = &products
		for _, product := range *schema.Products {
			*orderModel.Products = append(*orderModel.Products, product_service.ProductModel{
				nil,
				product.OrderID,
				product.Name,
				product.Description,
				product.Price,
				product.Currency,
				product.ShopUrl,
				product.ShopName,
				product.ImageUrls,
				nil,
				nil,
			})
		}
	}

	return orderModel
}
