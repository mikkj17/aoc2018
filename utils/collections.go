package utils

func Counter[T comparable](l []T) map[T]int {
	counts := map[T]int{}
	for _, e := range l {
		counts[e]++
	}

	return counts
}

func ContainsValue[T comparable, V comparable](m map[T]V, v V) bool {
	for _, value := range m {
		if value == v {
			return true
		}
	}

	return false
}
