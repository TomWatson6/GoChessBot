package linq

import "fmt"

func Any[T any](xs []T, f func(T) bool) bool {
	for _, x := range xs {
		if f(x) {
			return true
		}
	}

	return false
}

func All[T any](xs []T, f func(T) bool) bool {
	for _, x := range xs {
		if !f(x) {
			return false
		}
	}

	return true
}

func Where[T any](xs []T, f func(T) bool) []T {
	var result []T

	for _, x := range xs {
		if f(x) {
			result = append(result, x)
		}
	}

	return result
}

func Select[T any, R any](xs []T, f func(T) R) []R {
	var result []R

	for _, x := range xs {
		result = append(result, f(x))
	}

	return result
}

func Find[T any](xs []T, f func(T) bool) (*T, error) {
	for _, x := range xs {
		if f(x) {
			return &x, nil
		}
	}

	return nil, fmt.Errorf("couldn't locate a piece with condition specified")
}
