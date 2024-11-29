package gengoutil

import (
	"encoding/json"
	"errors"
	"fmt"
)

// check whether map has any key or not
func HasAnyKey[K comparable, V any](m map[K]V, keys ...K) bool {
	for _, key := range keys {
		if _, ok := m[key]; ok {
			return true
		}
	}
	return false
}

// coalesce value for map
func CoalesceMap[V any](m map[string]V, key string, def V) V {
	value, ok := m[key]
	if !ok {
		return def
	}

	if NilEmpty(&value) == nil {
		return def
	}

	return value
}

// convert map[string]any i.e. interface into map[string]type; will skip the key if type not match
func MapAnyToType[T any](input map[string]interface{}) map[string]T {
	output := make(map[string]T)
	for key, value := range input {
		// Attempt to convert the value to type T using type assertion
		if v, ok := value.(T); ok {
			output[key] = v
		} else {
			// skip the key if the type assertion fails.
			fmt.Printf("Skipping key %s: value cannot be converted to %T\n", key, value)
		}
	}
	return output
}

// convert struct into map[string]any
func StructToMapAny(i any) (map[string]any, error) {
	data, jsonErr := json.Marshal(i)
	if jsonErr != nil {
		return nil, errors.New(jsonErr.Error())
	}
	var ni map[string]any
	if jsonErr = json.Unmarshal(data, &ni); jsonErr != nil {
		return nil, errors.New(jsonErr.Error())
	}
	return ni, nil
}
