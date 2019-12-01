package godash

func IntFilter(elements []int, f func(e int) bool) []int {
	var filtered []int
	for _, element := range elements {
		if f(element) {
			filtered = append(filtered, element)
		}
	}
	return filtered
}
