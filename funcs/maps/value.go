package maps

// GetOr returns the value of the key in the map if it exists, otherwise it returns the default value.
func GetOr[T any](m map[string]any, key string, defval T) T {
	if v, ok := m[key]; ok {
		if val, ok := v.(T); ok {
			return val
		}
	}
	return defval
}
