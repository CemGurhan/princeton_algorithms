package stacksandqueues

import (
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

func (s *Stack[T]) Pop() T {
	length := len(*s)
	var item T

	if s.IsEmpty() {
		return item
	}

	item = (*s)[length-1]
	*s = (*s)[:length-1]

	return item
}

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

		comparionsItem := comparison.Pop()

		if item > comparionsItem {
			comparison = Stack[T]{}
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