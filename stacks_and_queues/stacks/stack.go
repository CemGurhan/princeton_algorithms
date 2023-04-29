package stacksandqueues

import (
	"errors"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	if len(*s) == 0 {
		return true
	}

	return false
}

func (s *Stack) Push(item string) *Stack {
	*s = append(*s, item)
	return s
}

func (s *Stack) Pop() (string, error) {
	length := len(*s)

	if s.IsEmpty() {
		return "", errors.New("stack empty")
	}

	item := (*s)[length-1]
	*s = (*s)[:length-1]

	return item, nil
}
