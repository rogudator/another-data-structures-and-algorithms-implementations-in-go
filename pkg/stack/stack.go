package stack

import "errors"

type Stack struct {
	items []int
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("Empty Stack")
	}
	l := len(s.items)
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove, nil
}