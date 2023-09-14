package utils

type Number interface {
	int | float32 | float64
}

func Sum[T Number](nums []T) T {
	var sum T
	for _, num := range nums {
		sum += num
	}
	return sum
}
