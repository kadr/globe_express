package travaler_domain_valueobjects

import (
	api_errors "github.com/kadr/globe_express/pkg/errors"
)

var (
	Planned   string = "planned"
	Ongoing   string = "ongoing"
	Completed string = "completed"
)

func NewStatus(status string) (string, error) {
	if status == "" {
		return Planned, nil
	}
	if status != Planned || status != Ongoing || status != Completed {
		return "", api_errors.ErrorCanNotCancel
	}
	return status, nil
}
