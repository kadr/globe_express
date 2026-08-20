package order_domain_order_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	vo "github.com/kadr/globe_express/internal/order_service/domain/value_objects"
	services "github.com/kadr/globe_express/internal/order_service/interfaces/api/handlers"
	api_errors "github.com/kadr/globe_express/pkg/errors"
)

func (os *OrderService) Cancel(ctx context.Context, orderID uuid.UUID) (services.OrderModel, error) {
	if order, err := os.repo.GetDetail(ctx, orderID); err == nil {
		if order.Status != vo.Pending {
			return services.OrderModel{}, fmt.Errorf("Can't cancel active order. %w", api_errors.ErrorCanNotCancel)
		}
	}
	order, err := os.repo.Cancel(ctx, orderID)
	if err != nil {
		return services.OrderModel{}, fmt.Errorf("order Cancel service error: %w", err)
	}

	return FromDomainModel(order), nil
}
