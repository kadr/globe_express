package order_api_order_dto

import (
	product_dto "github.com/kadr/globe_express/internal/order_service/application/dto/product"
	services "github.com/kadr/globe_express/internal/order_service/interfaces/api/handlers"
)

func ToDTO(schema services.OrderModel) OrderDTO {
	orderDTO := OrderDTO{
		schema.ID.String(),
		schema.CustomerID.String(),
		nil,
		*schema.Status,
		schema.PickupCity,
		schema.PickupCountry,
		schema.DeliveryCity,
		schema.DeliveryCountry,
		schema.RewardAmount,
		schema.RewardCurrency,
		schema.DeliveryDeadline,
		schema.DeliveryDate,
		schema.CreatedAt.Local().String(),
		nil,
		nil,
	}
	if schema.TravelerID != nil {
		travelerID := *schema.TravelerID
		travelerIDStr := travelerID.String()
		orderDTO.TravelerID = &travelerIDStr
	}
	if schema.UpdatedAt != nil {
		updateAt := *schema.UpdatedAt
		strUpdateAt := updateAt.Local().String()
		orderDTO.UpdatedAt = &strUpdateAt
	}
	if schema.Products != nil {
		products := make([]product_dto.ProductDTO, 0, len(*schema.Products))
		for _, product := range *schema.Products {
			item := product_dto.ProductDTO{
				product.ID.String(),
				nil,
				product.Name,
				product.Description,
				product.Price,
				product.Currency,
				product.ShopUrl,
				product.ShopName,
				product.ImageUrls,
				product.CreatedAt.String(),
				nil,
			}
			if product.OrderID != nil {
				orderID := product.OrderID.String()
				item.OrderID = &orderID
			}
			if product.UpdatedAt != nil {
				updateAt := product.UpdatedAt.String()
				item.UpdatedAt = &updateAt
			}
			products = append(products, item)
		}
		orderDTO.Products = &products
	}

	return orderDTO
}

func ToDTOList(schema []services.OrderModel) []OrderDTO {
	var results []OrderDTO
	for _, order := range schema {
		results = append(results, ToDTO(order))
	}

	return results
}
