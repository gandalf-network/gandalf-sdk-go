package graphqlTypes

import (
	"encoding/json"
	"time"
)

type Date time.Time

const dateFormat = "01/02/2006"

// UnmarshalJSON converts an input JSON string to a Date
func (d *Date) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	t, err := time.Parse(dateFormat, str)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

// MarshalJSON converts the Date to a JSON string
func (d Date) MarshalJSON() ([]byte, error) {
	dateStr := time.Time(d).Format(dateFormat)
	return json.Marshal(dateStr)
}
