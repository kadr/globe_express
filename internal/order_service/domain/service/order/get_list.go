package order_domain_order_service

import (
	"context"
	"fmt"

	services "github.com/kadr/globe_express/internal/order_service/interfaces/api/handlers"
)

func (os *OrderService) GetList(ctx context.Context, limit, offset int) ([]services.OrderModel, error) {
	orders, err := os.repo.GetList(ctx, limit, offset)
	if err != nil {
		return []services.OrderModel{}, fmt.Errorf("order GetList service error: %w", err)
	}

	return FromDomainModelList(orders), nil
}
