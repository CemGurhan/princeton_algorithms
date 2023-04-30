package practicequiz

import (
	"golang.org/x/exp/constraints"
)

// Queue with two stacks. Implement a queue with two stacks so
// that each queue operations takes a constant amortized number of stack operations.

type Queue[T constraints.Ordered] []T

func (q *Queue[T]) Enqueue(item T) *Queue[T] {
	return q.push(item)
}

func (q *Queue[T]) Dequeue() T {
	reverseQueue := Queue[T]{}
	queueLength := len(*q)
	var item T

	if q.isEmpty() {
		return item
	}

	for i := 0; i < queueLength; i++ {
		originalItem := q.pop()
		reverseQueue.push(originalItem)
	}

	dequeued := reverseQueue.pop()
	reverseQueueLength := len(reverseQueue)

	for i := 0; i < reverseQueueLength; i++ {
		reverseItem := reverseQueue.pop()
		q.push(reverseItem)
	}

	return dequeued
}

func (s *Queue[T]) isEmpty() bool {
	if len(*s) == 0 {
		return true
	}

	return false
}

func (s *Queue[T]) push(item T) *Queue[T] {
	*s = append(*s, item)
	return s
}

func (s *Queue[T]) pop() T {
	length := len(*s)
	var item T

	if s.isEmpty() {
		return item
	}

	item = (*s)[length-1]
	*s = (*s)[:length-1]

	return item
}
