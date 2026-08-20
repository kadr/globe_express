package order_domain_order_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	vo "github.com/kadr/globe_express/internal/order_service/domain/value_objects"
	services "github.com/kadr/globe_express/internal/order_service/interfaces/api/handlers"
	api_errors "github.com/kadr/globe_express/pkg/errors"
)

func (os *OrderService) Deliver(ctx context.Context, orderID uuid.UUID) (services.OrderModel, error) {
	order, err := os.GetDetail(ctx, orderID)
	if err != nil {
		return services.OrderModel{}, err
	}
	if *order.Status == string(vo.Delivered) {
		return services.OrderModel{}, fmt.Errorf("order already delivered. %w", api_errors.ErrorBadRequest)
	}
	if *order.Status != string(vo.InTransit) {
		return services.OrderModel{}, fmt.Errorf("order not transit yet. Current status is: %s. %w", *order.Status, api_errors.ErrorBadRequest)
	}
	status := string(vo.Delivered)
	schema, err := ToUpdateDomainModel(services.OrderUpdateModel{Status: &status})
	if err != nil {
		return services.OrderModel{}, fmt.Errorf("order Deliver service error: %w", err)
	}
	updatesOrder, err := os.repo.Update(ctx, orderID, schema)
	if err != nil {
		return services.OrderModel{}, fmt.Errorf("order Deliver service error: %w", err)
	}

	return FromDomainModel(updatesOrder), nil
}
