package nilx

// NilZero returns the zero value of the type of the pointer v if v is nil, otherwise it returns the value of v.
func NilZero[T comparable](val *T) T {
	if val == nil {
		var zero T

		return zero
	}

	return *val
}
