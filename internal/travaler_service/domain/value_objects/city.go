package travaler_domain_valueobjects

import (
	"fmt"

	api_errors "github.com/kadr/globe_express/pkg/errors"
)

func NewCity(city string) (string, error) {
	if len(city) < 3 {
		return "", fmt.Errorf("city must contains more than 3 leters. %w", api_errors.ErrorFieldValidation)
	}
	if len(city) > 20 {
		return "", fmt.Errorf("city must contains less than 20 leters. %w", api_errors.ErrorFieldValidation)
	}
	return city, nil
}
