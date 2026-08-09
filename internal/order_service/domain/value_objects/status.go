package order_domain_valueobjects

import (
	"fmt"
	"strings"

	api_errors "github.com/kadr/globe_express/pkg/errors"
)

type OrderStatus string

const (
	Pending   OrderStatus = "pending"
	Accepted  OrderStatus = "accepted"
	Purchased OrderStatus = "purchased"
	InTransit OrderStatus = "in_transit"
	Delivered OrderStatus = "delivered"
	Cancelled OrderStatus = "cancelled"
)

func NewStatus(status string) (OrderStatus, error) {
	status = strings.TrimSpace(status)
	status = strings.ToLower(status)
	if len([]rune(status)) == 0 {
		return "", fmt.Errorf("order status must be present. %w", api_errors.ErrorFieldValidation)
	}
	newStatus := OrderStatus(status)
	if newStatus != Pending && newStatus != Accepted && newStatus != Purchased && newStatus != InTransit && newStatus != Delivered && newStatus != Cancelled {
		return "", fmt.Errorf("incorrect order status. Status may be: pending, accepted, purchased, in_transit, delivered, cancelled %w", api_errors.ErrorIncorrectStatus)
	}
	return newStatus, nil
}
