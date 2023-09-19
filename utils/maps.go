package utils

import "cmp"

type Mode string

const (
	min Mode = "min"
	max Mode = "max"
)

func ContainsValue[T comparable, V comparable](m map[T]V, v V) bool {
	for _, value := range m {
		if value == v {
			return true
		}
	}

	return false

}

func MinValue[K comparable, V cmp.Ordered](m map[K]V) (K, V) {
	return extremeValue(m, min)
}

func MaxValue[K comparable, V cmp.Ordered](m map[K]V) (K, V) {
	return extremeValue(m, max)
}

func extremeValue[K comparable, V cmp.Ordered](m map[K]V, mode Mode) (K, V) {
	if len(m) == 0 {
		panic("cannot find extreme value of empty map")
	}

	var extremeKey K
	var extremeValue V

	for k, v := range m {
		extremeKey = k
		extremeValue = v
		break
	}

	for k, v := range m {
		if mode == min {
			if v < extremeValue {
				extremeKey = k
				extremeValue = v
			}
		} else {
			if v > extremeValue {
				extremeKey = k
				extremeValue = v
			}
		}
	}

	return extremeKey, extremeValue
}
