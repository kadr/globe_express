package travaler_domain_valueobjects

import "errors"

func NewWeight(weight float64) (float64, error) {
	if weight > 100.0 {
		return 0.0, errors.New("overload, weight max 100 kg.")
	}
	return weight, nil
}
