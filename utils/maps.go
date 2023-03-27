package utils

func Keys[T comparable](m map[T]any) []T {
	var keys []T
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}
