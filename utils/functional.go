package utils

func Map[T, S any](l []T, f func(T) S) []S {
	mapped := make([]S, 0, len(l))

	for _, val := range l {
		mapped = append(mapped, f(val))
	}

	return mapped
}

func SumBy[N Number, T any](l []T, f func(T) N) N {
	var sum N
	for _, val := range l {
		sum += f(val)
	}
	return sum
}
