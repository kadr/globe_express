package order_domain_valueobjects

import (
	"fmt"
	"strings"

	api_errors "github.com/kadr/globe_express/pkg/errors"
)

func NewName(name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return "", fmt.Errorf("name can't be empty. %w", api_errors.ErrorFieldValidation)
	}
	if len([]rune(name)) < 3 {
		return "", fmt.Errorf("name must be grate than 3. %w", api_errors.ErrorFieldValidation)
	}
	if len([]rune(name)) > 256 {
		return "", fmt.Errorf("name must be less than 256. %w", api_errors.ErrorFieldValidation)
	}
	return name, nil
}
