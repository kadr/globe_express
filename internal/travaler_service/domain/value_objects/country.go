package travaler_domain_valueobjects

import "errors"

func NewCountry(country string) (string, error) {
	if len(country) < 2 {
		return "", errors.New("country must contains more than 5 leters")
	}
	if len(country) > 20 {
		return "", errors.New("country must contains less than 20 leters")
	}
	return country, nil
}
