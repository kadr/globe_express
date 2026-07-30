package travaler_domain_valueobjects

import (
	"fmt"

	api_errors "github.com/kadr/globe_express/pkg/errors"
)

func NewWeight(weight float64) (float64, error) {
	if weight > 100.0 {
		return 0.0, fmt.Errorf("overload, weight max 100 kg. %w", api_errors.ErrorFieldValidation)
	}
	return weight, nil
}
