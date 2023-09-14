package sets

func Contains[T comparable](s Set[T], e T) bool {
	_, ok := s.Values[e]
	return ok
}

func Add[T comparable](s Set[T], e T) {
	s.Values[e] = struct{}{}
}

func Remove[T comparable](s Set[T], e T) {
	delete(s.Values, e)
}

func Union[T comparable](s Set[T], t Set[T]) Set[T] {
	u := Set[T]{}
	for e := range s.Values {
		Add(u, e)
	}
	for e := range t.Values {
		Add(u, e)
	}

	return u
}

func Intersection[T comparable](s Set[T], t Set[T]) Set[T] {
	v := Set[T]{}
	for e := range s.Values {
		if Contains(t, e) {
			Add(v, e)
		}
	}

	return v
}
