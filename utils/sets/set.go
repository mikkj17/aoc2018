package sets

type Set[T comparable] struct {
	Values map[T]struct{}
}

func Empty[T comparable]() Set[T] {
	return Set[T]{Values: map[T]struct{}{}}
}

func FromSlice[T comparable](vals []T) Set[T] {
	s := Set[T]{Values: map[T]struct{}{}}
	for _, val := range vals {
		Add(s, val)
	}

	return s
}

func ToSlice[T comparable](s Set[T]) []T {
	ret := make([]T, 0, Length(s))
	for e := range s.Values {
		ret = append(ret, e)
	}

	return ret
}
