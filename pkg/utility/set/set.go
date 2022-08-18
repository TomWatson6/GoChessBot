package set

type Set[T comparable] map[T]bool

func NewFromSlice[T comparable](items []T) Set[T] {
	var set Set[T]
	set = make(map[T]bool)

	set.AddMany(items)

	return set
}

func NewFromMapKeys[T comparable, R comparable](kvp map[T]R) Set[T] {
	var set Set[T]
	set = make(map[T]bool)

	for k := range kvp {
		set.Add(k)
	}

	return set
}

func NewFromMapValues[T comparable, R comparable](kvp map[T]R) Set[R] {
	var set Set[R]
	set = make(map[R]bool)

	for _, v := range kvp {
		set.Add(v)
	}

	return set
}

func (s *Set[T]) Add(item T) {
	(*s)[item] = true
}

func (s *Set[T]) AddMany(items []T) {
	for _, item := range items {
		(*s)[item] = true
	}
}

func (s *Set[T]) Remove(item T) {
	(*s)[item] = false
}

func (s *Set[T]) RemoveMany(items []T) {
	for _, item := range items {
		(*s)[item] = false
	}
}

func (s Set[T]) Intersect(t Set[T]) Set[T] {
	var intersection Set[T]

	for elem := range s {
		if _, ok := t[elem]; ok {
			intersection[elem] = true
		}
	}

	return intersection
}

func (s Set[T]) ToArray() []T {
	var arr []T

	for item := range s {
		arr = append(arr, item)
	}

	return arr
}
