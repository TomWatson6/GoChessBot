package set

import (
	"errors"
	"fmt"
)

type OrderedSet[T comparable] map[int]T

func (s *OrderedSet[T]) Add(item T) {
	(*s)[len(*s)] = item
}

func (s *OrderedSet[T]) AddAtIndex(index int, item T) error {
	if index > len(*s) {
		return errors.New("index of element to add to set specified is too large, should not be larger than length of set")
	}

	for i := len(*s); i > index; i-- {
		(*s)[i] = (*s)[i-1]
	}

	(*s)[index] = item

	return nil
}

func (s *OrderedSet[T]) Remove(item T) error {
	if index, err := s.find(item); err != nil {
		return s.remove(index)
	} else {
		return fmt.Errorf("cannot remove item specified: %w", err)
	}
}

func (s *OrderedSet[T]) RemoveAtIndex(index int) error {
	return s.remove(index)
}

func (s OrderedSet[T]) find(item T) (int, error) {
	for k, v := range s {
		if v == item {
			return k, nil
		}
	}

	return -1, fmt.Errorf("cannot find element specified in set: %v", item)
}

func (s *OrderedSet[T]) remove(index int) error {
	if index >= len(*s) {
		return errors.New("out of bounds error, index to remove from set too large")
	}

	for i := index; i < len(*s)-1; i++ {
		(*s)[i] = (*s)[i+1]
	}

	delete(*s, len(*s)-1)

	return nil
}
