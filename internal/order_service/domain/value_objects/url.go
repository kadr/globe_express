package order_domain_valueobjects

import (
	"fmt"
	"strings"

	api_errors "github.com/kadr/globe_express/pkg/errors"
)

func NewUrl(url string) (string, error) {
	url = strings.TrimSpace(url)
	if len(url) == 0 {
		return "", fmt.Errorf("url can't be empty. %w", api_errors.ErrorFieldValidation)
	}
	if !strings.Contains(url, "https://") && !strings.Contains(url, "http://") {
		return "", fmt.Errorf("url must countain http or https. %w", api_errors.ErrorFieldValidation)
	}
	return url, nil
}
