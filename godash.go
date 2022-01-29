// Package godash provide lodash like array manipulation utility functions.
package godash

// Filter accepts array and functions to filter and returns array filtered by given function.
func Filter[T any](input []T, filterFunc func(e T) bool) []T {
	var filtered []T
	for _, v := range input {
		if filterFunc(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

// Map accepts array and functions to map and returns array filtered by given function.
func Map[T any, K any](input []T, mapFunc func(e T) K) []K {
	var mapped []K
	for _, v := range input {
		mapped = append(mapped, mapFunc(v))
	}
	return mapped
}

func Reverse[T any](input []T) []T {
	ilen := len(input)
	reversed := make([]T, ilen, ilen)
	for idx, v := range input {
		reversed[ilen-1-idx] = v
	}
	return reversed
}
