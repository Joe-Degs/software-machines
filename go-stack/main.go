package main

import "fmt"

type stack_status int

const (
	STACK_OK = stack_status(iota)
	STACK_OVERFLOW
	STACK_UNDERFLOW

	STACK_MAX = 100
)

type stack struct {
	data []int
}

func (s *stack) Push(data int) stack_status {
	if len(s.data) > STACK_MAX {
		return STACK_OVERFLOW
	}
	s.data = append(s.data, data)
	return STACK_OK
}

func (s *stack) Pop() (int, stack_status) {
	if s == nil || len(s.data) < 1 {
		return 0, STACK_UNDERFLOW
	}
	sp := len(s.data) - 1
	data := s.data[sp]
	s.data = s.data[:sp]
	return data, STACK_OK
}

func (s *stack) Depth() int {
	return len(s.data)
}

func main() {
	s := stack{}
	s.Push(1)
	s.Push(2)
	fmt.Printf("depth = %d\n", s.Depth())
	l, _ := s.Pop()
	r, _ := s.Pop()
	fmt.Printf("l = %d\n", l)
	fmt.Printf("d = %d\n", r)
	fmt.Printf("depth = %d\n", s.Depth())
}
