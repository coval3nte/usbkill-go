package utils

func DeepCopy[T comparable, Z comparable](dst, src map[T]Z) {
	for k, v := range src {
		dst[k] = v
	}
}

func MapKeys[T comparable, Z interface{}](t map[T]Z) []T {
	keys := make([]T, 0, len(t))
	for k := range t {
		keys = append(keys, k)
	}
	return keys
}
