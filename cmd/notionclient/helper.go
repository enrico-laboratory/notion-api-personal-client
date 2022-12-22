package notionclient

import (
	"errors"
	"time"
)

func ConvertNotionDateTime(taskDate string) (*time.Time, error) {
	if len(taskDate) == 10 {
		parsedDate, err := time.Parse("2006-01-02", taskDate)
		if err != nil {
			return nil, err
		}
		return &parsedDate, nil
	}
	if len(taskDate) < 10 {
		return nil, errors.New("invalid date length")
	}
	if len(taskDate) > 10 {
		parsedDate, err := time.Parse(time.RFC3339, taskDate)
		if err != nil {
			return nil, err
		}
		return &parsedDate, nil
	}
	return nil, errors.New("something is wrong in the date/time format")
}
