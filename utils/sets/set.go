package sets

type Set[T comparable] struct {
	Values map[T]struct{}
}

func New[T comparable]() Set[T] {
	return Set[T]{map[T]struct{}{}}
}
