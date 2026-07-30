package order_domain_valueobjects

import (
	"fmt"
	"strings"

	api_errors "github.com/kadr/globe_express/pkg/errors"
)

func NewDescription(description string) (string, error) {
	description = strings.TrimSpace(description)
	if description == "" {
		return "", fmt.Errorf("description can't be empty. %w", api_errors.ErrorFieldValidation)
	}
	if len([]rune(description)) < 3 {
		return "", fmt.Errorf("description must be grate than 3 characters. %w", api_errors.ErrorFieldValidation)
	}
	if len([]rune(description)) > 1000 {
		return "", fmt.Errorf("description must be less than 1000 characters. %w", api_errors.ErrorFieldValidation)
	}
	return description, nil
}
