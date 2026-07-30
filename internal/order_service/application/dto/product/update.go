package order_api_product_dto

import (
	service "github.com/kadr/globe_express/internal/order_service/domain/service/product"
)

type UpdateDTO struct {
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
	Price       *float64  `json:"price"`
	Currency    *string   `json:"currency"`
	ShopUrl     *string   `json:"shop_url"`
	ShopName    *string   `json:"shop_name"`
	ImageUrls   *[]string `json:"image_urls"`
}

func ToUpdateModel(schema UpdateDTO) service.ProductUpdateModel {
	productModel := service.ProductUpdateModel{
		schema.Name,
		schema.Description,
		schema.Price,
		schema.Currency,
		schema.ShopUrl,
		schema.ShopName,
		schema.ImageUrls,
	}

	return productModel
}
