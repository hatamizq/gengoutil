package gengoutil

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
