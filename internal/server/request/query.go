package request

import (
	"fmt"
	"strconv"
)

type MissingRequiredQueryParameterError struct {
	Parameter string
}

func (m *MissingRequiredQueryParameterError) Error() string {
	return fmt.Sprintf("missing required query parameter %s", m.Parameter)
}

func (r *Request) GetQueryValue(required bool, key, defaultValue string) (string, error) {
	q := r.Request.URL.Query()
	if !q.Has(key) {
		if !required {
			return defaultValue, nil
		}
		return defaultValue, &MissingRequiredQueryParameterError{Parameter: key}
	}
	return q.Get(key), nil
}

func (r *Request) GetQueryValueInt(required bool, key string, defaultValue, min int) (int, error) {
	v, err := r.GetQueryValue(required, key, "")
	if err != nil {
		return 0, err
	}
	if v == "" {
		if !required {
			return defaultValue, nil
		}
		return 0, &MissingRequiredQueryParameterError{Parameter: key}
	}
	ans, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("invalid value for query parameter %s: %v", key, err)
	}
	if ans < min {
		return 0, fmt.Errorf("invalid value for query parameter %s: %d is less than %d", key, ans, min)
	}
	return ans, nil
}
