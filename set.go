package set

type Set[T comparable] struct {
	hash map[T]struct{}
}

// Create a new set
func New[T comparable](initial ...T) *Set[T] {
	s := &Set[T]{make(map[T]struct{})}

	for _, v := range initial {
		s.Insert(v)
	}

	return s
}

func (s *Set[T]) Insert(value T) {
	s.hash[value] = struct{}{}
}

func (s *Set[T]) Remove(value T) {
	delete(s.hash, value)
}

func (s *Set[T]) Has(value T) bool {
	_, exists := s.hash[value]
	return exists
}

func (s *Set[T]) Len() int {
	return len(s.hash)
}

func (s *Set[T]) Elements() []T {
	elements := make([]T, len(s.hash))
	i := 0
	for key := range s.hash {
		elements[i] = key
		i++
	}
	return elements
}

func (s *Set[T]) Intersection(t *Set[T]) *Set[T] {
	var a *Set[T]
	var b *Set[T]

	if s.Len() > t.Len() {
		a = t
		b = s
	} else {
		a = s
		b = t
	}

	intersection := New[T]()
	for key := range a.hash {
		if b.Has(key) {
			intersection.Insert(key)
		}
	}
	return intersection
}

func (s *Set[T]) Difference(t *Set[T]) *Set[T] {
	difference := New[T]()

	for key := range s.hash {
		if !t.Has(key) {
			difference.Insert(key)
		}
	}

	return difference
}

func (s *Set[T]) Union(t *Set[T]) *Set[T] {
	var a *Set[T]
	var b *Set[T]

	if s.Len() > t.Len() {
		a = t
		b = s
	} else {
		a = s
		b = t
	}

	union := b

	for key := range a.hash {
		union.Insert(key)
	}

	return union
}

func (s *Set[T]) SubsetOf(t *Set[T]) bool {

	if s.Len() > t.Len() {
		return false
	}

	for key := range s.hash {
		if !t.Has(key) {
			return false
		}
	}

	return true
}

func (s *Set[T]) ProperSubsetOf(t *Set[T]) bool {
	if s.Len() == t.Len() {
		return false
	}
	return s.SubsetOf(t)
}

func (s *Set[T]) Copy() *Set[T] {
	copy := New[T]()

	for key := range s.hash {
		copy.Insert(key)
	}

	return copy
}

func (s *Set[T]) Equals(t *Set[T]) bool {
	if s.Len() != t.Len() {
		return false
	}

	for key := range s.hash {
		if !t.Has(key) {
			return false
		}
	}

	return true
}
