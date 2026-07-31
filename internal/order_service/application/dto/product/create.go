package order_api_product_dto

import (
	"github.com/google/uuid"
	service "github.com/kadr/globe_express/internal/order_service/domain/service/product"
)

type CreateDTO struct {
	OrderID     *uuid.UUID `json:"order_id"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	Price       float64    `json:"price"`
	Currency    string     `json:"currency"`
	ShopUrl     string     `json:"shop_url"`
	ShopName    string     `json:"shop_name"`
	ImageUrls   *[]string  `json:"image_urls"`
}

func ToCreateModel(schema CreateDTO) service.ProductModel {
	productModel := service.ProductModel{
		nil,
		schema.OrderID,
		schema.Name,
		schema.Description,
		schema.Price,
		schema.Currency,
		schema.ShopUrl,
		schema.ShopName,
		schema.ImageUrls,
		nil,
		nil,
	}

	return productModel
}
