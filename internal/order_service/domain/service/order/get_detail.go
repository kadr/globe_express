package order_domain_order_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	services "github.com/kadr/globe_express/internal/order_service/interfaces/api/handlers"
)

func (os *OrderService) GetDetail(ctx context.Context, orderID uuid.UUID) (services.OrderModel, error) {
	order, err := os.repo.GetDetail(ctx, orderID)
	if err != nil {
		return services.OrderModel{}, fmt.Errorf("order GetDetail service error: %w", err)
	}

	return FromDomainModel(order), nil
}
