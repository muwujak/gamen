package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JSON map[string]interface{}

func (j JSON) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}

	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("cannot scan %T into JSON, accepted values are []byte and string", value)
	}

	return json.Unmarshal(bytes, j)
}

func (j *JSON) Transform(value any) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = j.Scan(bytes)
	if err != nil {
		return err
	}

	return nil
}
