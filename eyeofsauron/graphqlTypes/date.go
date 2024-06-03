package graphqlTypes

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Date time.Time

const dateFormat = "01/02/2006"

// UnmarshalJSON converts a JSON string to a Date
func (d *Date) UnmarshalJSON(data []byte) error {
	str, err := strconv.Unquote(string(data))
	if err != nil {
		return fmt.Errorf("date must be a string: %w", err)
	}

	t, err := time.Parse(dateFormat, str)
	if err != nil {
		return fmt.Errorf("invalid date format: %w", err)
	}

	*d = Date(t)
	return nil
}

// MarshalJSON converts the Date to a JSON string
func (d Date) MarshalJSON() ([]byte, error) {
	dateStr := time.Time(d).Format(dateFormat)

	// Quote the string and return as JSON
	return json.Marshal(dateStr)
}
