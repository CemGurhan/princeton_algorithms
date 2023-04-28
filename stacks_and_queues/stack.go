package stacksandqueues

type stack []string

func (s stack) IsEmpty() bool {
	if len(s) == 0 {
		return true
	}

	return false
}

func (s stack) Push(item string) stack {
	s = append(s, item)
	return s
}

func (s stack) Pop() string {
	length := len(s)

	if s.IsEmpty() {
		return ""
	}

	item := s[length-1]
	s = s[:length-1]

	return item
}
