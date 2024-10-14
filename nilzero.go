package nilx

// NilZero returns the zero value of the type of the pointer val if val is nil, otherwise it returns the value of val.
func NilZero[T comparable](val *T) T {
	if val == nil {
		var zero T

		return zero
	}

	return *val
}
