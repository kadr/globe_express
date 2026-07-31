package order_api_product_dto

import (
	product_service "github.com/kadr/globe_express/internal/order_service/domain/service/product"
)

func ToDTO(schema product_service.ProductModel) ProductDTO {
	productDTO := ProductDTO{
		schema.ID.String(),
		nil,
		schema.Name,
		schema.Description,
		schema.Price,
		schema.Currency,
		schema.ShopUrl,
		schema.ShopName,
		schema.ImageUrls,
		schema.CreatedAt.Local().String(),
		nil,
	}
	if schema.OrderID != nil {
		orderID := *schema.OrderID
		orderIDStr := orderID.String()
		productDTO.UpdatedAt = &orderIDStr
	}
	if schema.UpdatedAt != nil {
		updateAt := *schema.UpdatedAt
		strUpdateAt := updateAt.Local().String()
		productDTO.UpdatedAt = &strUpdateAt
	}

	return productDTO
}

func ToDTOList(schema []product_service.ProductModel) []ProductDTO {
	var results []ProductDTO
	for _, product := range schema {
		results = append(results, ToDTO(product))
	}

	return results
}
