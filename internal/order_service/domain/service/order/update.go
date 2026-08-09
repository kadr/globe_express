package order_domain_order_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	domain_models "github.com/kadr/globe_express/internal/order_service/domain/models"
	services "github.com/kadr/globe_express/internal/order_service/interfaces/api/handlers"
)

func (os *OrderService) Update(ctx context.Context, orderID uuid.UUID, schema services.OrderUpdateModel) (services.OrderModel, error) {
	newSchema, err := ToUpdateDomainModel(schema)
	if err != nil {
		return services.OrderModel{}, fmt.Errorf("order update service error: %w", err)
	}
	order, err := os.repo.Update(ctx, orderID, newSchema)
	if err != nil {
		return services.OrderModel{}, fmt.Errorf("order update service error: %w", err)
	}

	return FromDomainModel(order), nil
}

func ToUpdateDomainModel(schema services.OrderUpdateModel) (domain_models.OrderUpdateModel, error) {
	orderModel, err := domain_models.NewUpdateOrder(
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
		return domain_models.OrderUpdateModel{}, err
	}

	return orderModel, nil
}
