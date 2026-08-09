package order_domain_valueobjects

import (
	"fmt"
	"strings"

	api_errors "github.com/kadr/globe_express/pkg/errors"
)

func NewLocation(location string) (string, error) {
	location = strings.TrimSpace(location)
	if location == "" {
		return "", fmt.Errorf("location can't be empty. %w", api_errors.ErrorFieldValidation)
	}
	if len([]rune(location)) < 2 {
		return "", fmt.Errorf("location must be grate or equal than 2. %w", api_errors.ErrorFieldValidation)
	}
	if len([]rune(location)) > 256 {
		return "", fmt.Errorf("location must be less than 256. %w", api_errors.ErrorFieldValidation)
	}
	return location, nil
}
