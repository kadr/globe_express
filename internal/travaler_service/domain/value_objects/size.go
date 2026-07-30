package travaler_domain_valueobjects

import (
	"fmt"

	api_errors "github.com/kadr/globe_express/pkg/errors"
)

func NewSize(size float64) (float64, error) {
	if size > 50.0 {
		return 0.0, fmt.Errorf("oversize, size max 50 cm3. %w", api_errors.ErrorFieldValidation)
	}
	return size, nil
}
