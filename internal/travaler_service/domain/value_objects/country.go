package travaler_domain_valueobjects

import (
	"fmt"

	api_errors "github.com/kadr/globe_express/pkg/errors"
)

func NewCountry(country string) (string, error) {
	if len(country) < 2 {
		return "", fmt.Errorf("country must contains more than 5 leters. %w", api_errors.ErrorFieldValidation)
	}
	if len(country) > 20 {
		return "", fmt.Errorf("country must contains less than 20 leters. %w", api_errors.ErrorFieldValidation)
	}
	return country, nil
}
