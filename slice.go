package gengoutil

import (
	"errors"
	"reflect"
	"sort"
)

// to determine whether slice contains value or not
func SliceContains[S ~[]E, E comparable](s S, v E) bool {
	return SliceIndex(s, v) >= 0
}

// to get index number of slice
func SliceIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// data type for sort ordering
type SortOrder string

const (
	AscendingSort  SortOrder = "asc"
	DescendingSort SortOrder = "desc"
)

// to sort slice struct by key; following with order
func SortSliceStructByKey[T any](slice []T, key string, order SortOrder) error {
	if len(slice) == 0 {
		return errors.New("slice cannot be empty to sort")
	}
	elemType := reflect.TypeOf(slice[0])
	if elemType.Kind() != reflect.Struct {
		return errors.New("slice element must be a struct")
	}
	sort.Slice(slice, func(i, j int) bool {
		valI := reflect.ValueOf(slice[i])
		valJ := reflect.ValueOf(slice[j])
		fieldI := valI.FieldByName(key)
		fieldJ := valJ.FieldByName(key)
		if !fieldI.IsValid() || !fieldJ.IsValid() {
			return false
		}
		return compareValues(fieldI, fieldJ, order)
	})
	return nil
}

// to sort slice; following with order
func SortSlice[T any](slice []T, order SortOrder) error {
	if len(slice) == 0 {
		return errors.New("slice cannot be empty to sort")
	}

	sort.Slice(slice, func(i, j int) bool {
		valI := reflect.ValueOf(slice[i])
		valJ := reflect.ValueOf(slice[j])

		return compareValues(valI, valJ, order)
	})
	return nil
}

// compareValues compares two reflect.Values and returns true if the first value
// should come before the second based on the order ("asc" or "desc").
func compareValues(fieldI, fieldJ reflect.Value, order SortOrder) bool {
	if fieldI.Type() != fieldJ.Type() {
		return false
	}

	switch fieldI.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		if order == AscendingSort {
			return fieldI.Int() < fieldJ.Int()
		}
		return fieldI.Int() > fieldJ.Int()
	case reflect.Float32, reflect.Float64:
		if order == AscendingSort {
			return fieldI.Float() < fieldJ.Float()
		}
		return fieldI.Float() > fieldJ.Float()
	case reflect.String:
		if order == AscendingSort {
			return fieldI.String() < fieldJ.String()
		}
		return fieldI.String() > fieldJ.String()
	default:
		return false
	}
}
