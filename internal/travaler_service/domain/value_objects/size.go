package travaler_domain_valueobjects

import "errors"

func NewSize(size float64) (float64, error) {
	if size > 50.0 {
		return 0.0, errors.New("oversize, size max 50 cm3")
	}
	return size, nil
}
