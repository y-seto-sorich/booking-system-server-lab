package util

import (
	"fmt"
	"time"
)

func ToTime(dateStr string, layout string) (time.Time, error) {
	parsedTime, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format: %w", err)
	}
	return parsedTime, nil

	// t, err := time.ParseInLocation("2006-01-02", inDate, time.Local)
	// if err != nil {
	// 	return nil, errors.New("invalid date format for day")
	// }

	// outDate = Pointer(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local))
}

func ToDateString(t time.Time) string {
	return t.Format("2006-01-02")
}

func Now() time.Time {
	return time.Now()
}
