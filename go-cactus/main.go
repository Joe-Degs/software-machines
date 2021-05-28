package main

import "fmt"

// cactus stack is a singly linked list.
type stack struct {
	data int
	prev *stack
}

// methods push, pop, depth

// add new item to top of stack
// return that item
func (s *stack) push(data int) *stack {
	ns := &stack{data: data, prev: s}
	return ns
}

func (s *stack) pop() (*stack, int) {
	return s.prev, s.data
}

func (s *stack) depth() int {
	var r int
	for t := s; t != nil; t = t.prev {
		r += 1
	}
	return r
}

func (s *stack) sum() int {
	var r int
	for t := s; t != nil; t = t.prev {
		r += t.data
	}
	return r
}

func main() {
	s := &stack{}
	s = s.push(3)
	s = s.push(4).push(10)
	fmt.Println("depth =", s.depth())
	fmt.Println("sum =", s.sum())
	s, d := s.pop()
	fmt.Printf("stack depth = %d, popped value = %d\n", s.depth(), d)
}
