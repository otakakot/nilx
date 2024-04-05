package nilex

// NilDef returns the value of the pointer v if v is not nil, otherwise it returns the default value def.
func NilDef[T any](v *T, def T) T {
	if v == nil {
		return def
	}

	return *v
}
