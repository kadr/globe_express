package travaler_domain_valueobjects

import (
	"fmt"
	"strings"

	api_errors "github.com/kadr/globe_express/pkg/errors"
)

var (
	Planned   string = "planned"
	Ongoing   string = "ongoing"
	Completed string = "completed"
)

func NewStatus(status string) (string, error) {
	status = strings.TrimSpace(status)
	if status == "" {
		return Planned, nil
	}
	if status != Planned && status != Ongoing && status != Completed {
		return "", fmt.Errorf("status may be only: planned, ongoing, completed. %w", api_errors.ErrorIncorrectStatus)
	}
	return status, nil
}
