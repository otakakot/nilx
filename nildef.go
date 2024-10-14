package nilx

// NilDef returns the value of the pointer val if val is not nil, otherwise it returns the default value def.
func NilDef[T comparable](val *T, def T) T {
	if val == nil {
		return def
	}

	return *val
}
