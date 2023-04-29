package practicequiz

import (
	"errors"

	sq "github.com/cemgurhan/princetonalgo/stacks_and_queues/stacks"
)

// Queue with two stacks. Implement a queue with two stacks so
// that each queue operations takes a constant amortized number of stack operations.

func Enqueue(item string, queue sq.Stack[string]) *sq.Stack[string] {
	return queue.Push(item)
}

func Dequeue(queue sq.Stack[string]) (string, error) {
	reverseQueue := sq.Stack[string]{}

	if queue.IsEmpty() {
		return "", errors.New("queue empty")
	}

	for _, item := range queue {
		reverseQueue.Push(item)
	}

	dequeued, _ := reverseQueue.Pop()

	return dequeued, nil
}
