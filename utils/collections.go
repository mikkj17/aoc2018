package utils

func Counter[T comparable](l []T) map[T]int {
	counts := map[T]int{}
	for _, e := range l {
		counts[e]++
	}

	return counts
}
