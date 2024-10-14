package nilx

// Or returns the first non-nil value from vals or the zero value of T if all values are nil.
func Or[T comparable](vals ...*T) T {
	var zero T

	for _, val := range vals {
		if val != nil {
			return *val
		}
	}

	return zero
}
