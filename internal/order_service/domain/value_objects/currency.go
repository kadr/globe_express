package order_domain_valueobjects

import (
	"fmt"
	"strings"

	api_errors "github.com/kadr/globe_express/pkg/errors"
)

const (
	Usd string = "USD"
	Rur string = "RUR"
	Eur string = "EUR"
)

func NewCurrency(currency string) (string, error) {
	currency = strings.TrimSpace(currency)
	currency = strings.ToUpper(currency)
	if len(currency) == 0 {
		return "", fmt.Errorf("currency can't be empty. %w", api_errors.ErrorFieldValidation)
	}
	if currency != Usd && currency != Rur && currency != Eur {
		return "", fmt.Errorf("currency can be one of: USD, RUR, EUR. %w", api_errors.ErrorFieldValidation)
	}
	return currency, nil
}
