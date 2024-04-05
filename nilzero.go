package nilex

import "reflect"

// NilZero returns the zero value of the type of the pointer v if v is nil, otherwise it returns the value of v.
func NilZero[T any](v *T) T {
	if v == nil {
		var zero T

		return zero
	}

	return *v
}

func IsZero[T any](v *T) bool {
	t := reflect.TypeOf(v)

	switch t.Kind() {
	case reflect.Ptr, reflect.Interface:
		return v == nil
	case reflect.Slice, reflect.Map, reflect.Array:
		return reflect.ValueOf(v).Len() == 0
	default:
		return reflect.DeepEqual(v, reflect.Zero(t).Interface())
	}
}
