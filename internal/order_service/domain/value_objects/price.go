package order_domain_valueobjects

import (
	"fmt"

	api_errors "github.com/kadr/globe_express/pkg/errors"
)

func NewPrice(price float64) (float64, error) {
	if price <= 0.0 {
		return 0.0, fmt.Errorf("price can't be negative or 0. %w", api_errors.ErrorFieldValidation)
	}
	return price, nil
}
