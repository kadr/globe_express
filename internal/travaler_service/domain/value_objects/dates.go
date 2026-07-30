package travaler_domain_valueobjects

import (
	"time"
)

func NewDate(date string) (time.Time, error) {
	newDate, err := time.Parse("2006-01-02T15:04:05Z07:00", date)
	if err != nil {
		return newDate, nil
	}
	return time.Time{}, err
}
