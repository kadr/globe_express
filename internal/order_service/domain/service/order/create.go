package order_domain_order_service

import (
	"context"
	"fmt"

	domain_models "github.com/kadr/globe_express/internal/order_service/domain/models"
	services "github.com/kadr/globe_express/internal/order_service/interfaces/api/handlers"
)

func (os *OrderService) Create(ctx context.Context, schema services.OrderModel) (services.OrderModel, error) {
	domainSchema, err := ToDomainModel(schema)
	if err != nil {
		return services.OrderModel{}, fmt.Errorf("order service error: %w", err)
	}
	order, err := os.repo.Create(ctx, domainSchema)
	if err != nil {
		return services.OrderModel{}, fmt.Errorf("order service error: %w", err)
	}
	if schema.Products != nil {
		for _, product := range *schema.Products {
			product.OrderID = &order.ID
			_, err := os.productService.Create(ctx, product)
			if err != nil {
				return FromDomainModel(order), fmt.Errorf("Can't create product. %w", err)
			}
		}
	}

	return FromDomainModel(order), nil
}

func ToDomainModel(schema services.OrderModel) (domain_models.OrderModel, error) {
	orderModel, err := domain_models.NewOrder(
		schema.CustomerID,
		schema.PickupCity,
		schema.PickupCountry,
		schema.DeliveryCity,
		schema.DeliveryCountry,
		schema.RewardCurrency,
		schema.RewardAmount,
		schema.DeliveryDate,
		schema.DeliveryDeadline,
		schema.Status,
		schema.TravelerID,
	)
	if err != nil {
		return domain_models.OrderModel{}, err
	}

	return orderModel, nil
}
