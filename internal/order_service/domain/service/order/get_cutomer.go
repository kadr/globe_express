package order_domain_order_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	services "github.com/kadr/globe_express/internal/order_service/interfaces/api/handlers"
)

func (os *OrderService) GetCustomer(ctx context.Context, customerID uuid.UUID, limit, offset int) ([]services.OrderModel, error) {
	orders, err := os.repo.GetCustomer(ctx, customerID, limit, offset)
	if err != nil {
		return []services.OrderModel{}, fmt.Errorf("order GetCustomer service error: %w", err)
	}

	return FromDomainModelList(orders), nil
}
