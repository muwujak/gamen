package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func GetSchema(iface any) (map[string]interface{}, error) {
	t := reflect.TypeOf(iface)
	schema := make(map[string]interface{})

	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("type is not a struct")
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		// Get the JSON tag, default to field name if no tag
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue // Skip fields without JSON tags or marked as ignored
		}
		// Use the first part of the JSON tag (ignore omitempty, etc.)
		jsonKey := jsonTag
		if commaIdx := strings.Index(jsonTag, ","); commaIdx != -1 {
			jsonKey = jsonTag[:commaIdx]
		}

		// Handle the field type
		fieldType := field.Type
		if fieldType.Kind() == reflect.Struct {
			// Recursively generate schema for nested struct
			nestedSchema, err := GetSchema(reflect.New(fieldType).Elem().Interface())
			if err != nil {
				return nil, fmt.Errorf("failed to get schema for field %s: %v", jsonKey, err)
			}
			schema[jsonKey] = nestedSchema
		} else if fieldType.Kind() == reflect.Slice {
			// Handle slices/arrays
			elemType := fieldType.Elem()
			if elemType.Kind() == reflect.Struct {
				// Recursively generate schema for the struct element type
				nestedSchema, err := GetSchema(reflect.New(elemType).Elem().Interface())
				if err != nil {
					return nil, fmt.Errorf("failed to get schema for slice element %s: %v", jsonKey, err)
				}
				schema[jsonKey] = map[string]interface{}{
					"type":  "array",
					"items": nestedSchema,
				}
			} else {
				// Non-struct slice (e.g., []string)
				schema[jsonKey] = map[string]interface{}{
					"type":  "array",
					"items": elemType.String(),
				}
			}
		} else {
			// Use the type's string representation for non-struct types
			schema[jsonKey] = fieldType.String()
		}
	}

	return schema, nil
}

func ChangeInto(source interface{}, dest interface{}) error {

	bytes, err := json.Marshal(source)
	if err != nil {
		return fmt.Errorf("failed to marshal configuration: %w", err)
	}

	err = json.Unmarshal(bytes, &dest)
	if err != nil {
		return fmt.Errorf("failed to unmarshal configuration: %w", err)
	}

	return nil
}
