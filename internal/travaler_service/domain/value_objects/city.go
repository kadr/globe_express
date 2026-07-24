package travaler_domain_valueobjects

import "errors"

func NewCity(city string) (string, error) {
	if len(city) < 3 {
		return "", errors.New("city must contains more than 5 leters")
	}
	if len(city) > 20 {
		return "", errors.New("city must contains less than 20 leters")
	}
	return city, nil
}
