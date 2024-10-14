package nilx

// Def returns the value of the pointer v if v is not nil, otherwise it returns the default value def.
func Def[T comparable](v *T, def T) T {
	if v == nil {
		return def
	}

	return *v
}
