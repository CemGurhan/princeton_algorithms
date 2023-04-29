package stacksandqueues

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type Stack[T constraints.Ordered] []T

func (s *Stack[T]) IsEmpty() bool {
	if len(*s) == 0 {
		return true
	}

	return false
}

func (s *Stack[T]) Push(item T) *Stack[T] {
	*s = append(*s, item)
	return s
}

func (s *Stack[T]) Pop() (T, error) {
	length := len(*s)
	var item T

	if s.IsEmpty() {
		return item, errors.New("stack empty")
	}

	item = (*s)[length-1]
	*s = (*s)[:length-1]

	return item, nil
}

func (s *Stack[T]) Clear() {
	for range *s {
		s.Pop()
	}
}

// TODO optimize
func (s *Stack[T]) FindMaxItem() *Stack[T] {
	comparison := Stack[T]{}

	if s.IsEmpty() {
		return &comparison
	}

	for _, item := range *s {
		if comparison.IsEmpty() {
			comparison.Push(item)
			continue
		}

		comparionsItem, _ := comparison.Pop()

		if item > comparionsItem {
			comparison.Clear()
			comparison.Push(item)
			continue
		} else if comparionsItem > item {
			comparison.Push(comparionsItem)
			continue
		}

		if item == comparionsItem {
			comparison.Push(item)
			comparison.Push(comparionsItem)
		}
	}

	return &comparison
}
