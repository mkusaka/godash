package godash

func Filter[T any](input []T, filterFunc func(e T) bool) []T {
	var filtered []T
	for _, v := range input {
		if filterFunc(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func Map[T any, K any](input []T, mapFunc func(e T) K) []K {
	var mapped []K
	for _, v := range input {
		mapped = append(mapped, mapFunc(v))
	}
	return mapped
}
