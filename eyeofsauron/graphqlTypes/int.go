package graphqlTypes

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Int int
type Int32 int32
type Int64 int64

func (i Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(i))
}

func (i *Int) UnmarshalJSON(data []byte) error {
	var v any
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v := v.(type) {
	case string:
		val, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*i = Int(val)
	case float64:
		*i = Int(v)
	case int:
		*i = Int(v)
	case int64:
		*i = Int(v)
	case json.Number:
		val, err := strconv.Atoi(v.String())
		if err != nil {
			return err
		}
		*i = Int(val)
	default:
		return fmt.Errorf("%T is not an int", v)
	}
	return nil
}

// Int32 - JSON Marshal and Unmarshal
func (i Int32) MarshalJSON() ([]byte, error) {
	return json.Marshal(int32(i))
}

func (i *Int32) UnmarshalJSON(data []byte) error {
	var v any
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v := v.(type) {
	case string:
		val, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return err
		}
		*i = Int32(val)
	case float64:
		*i = Int32(v)
	case int:
		*i = Int32(v)
	case int64:
		*i = Int32(v)
	case json.Number:
		val, err := strconv.ParseInt(v.String(), 10, 32)
		if err != nil {
			return err
		}
		*i = Int32(val)
	default:
		return fmt.Errorf("%T is not an int32", v)
	}
	return nil
}

// Int64 - JSON Marshal and Unmarshal
func (i Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(int64(i))
}

func (i *Int64) UnmarshalJSON(data []byte) error {
	var v any
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v := v.(type) {
	case string:
		val, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}
		*i = Int64(val)
	case float64:
		*i = Int64(v)
	case int:
		*i = Int64(v)
	case int64:
		*i = Int64(v)
	case json.Number:
		val, err := strconv.ParseInt(v.String(), 10, 64)
		if err != nil {
			return err
		}
		*i = Int64(val)
	default:
		return fmt.Errorf("%T is not an int64", v)
	}
	return nil
}
