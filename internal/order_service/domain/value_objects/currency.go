package order_domain_valueobjects

import (
	"fmt"
	"strings"

	shared_money "github.com/kadr/globe_express/internal/shared/money"
	api_errors "github.com/kadr/globe_express/pkg/errors"
)

func NewCurrency(currency string) (shared_money.Currency, error) {
	currency = strings.TrimSpace(currency)
	currency = strings.ToUpper(currency)
	if len(currency) == 0 {
		return "", fmt.Errorf("currency can't be empty. %w", api_errors.ErrorFieldValidation)
	}
	newCurrency := shared_money.Currency(currency)
	if newCurrency != shared_money.Usd && newCurrency != shared_money.Rub && newCurrency != shared_money.Eur {
		return "", fmt.Errorf("currency can be one of: USD, RUB, EUR. %w", api_errors.ErrorFieldValidation)
	}
	return newCurrency, nil
}
