package main

type Stack struct {
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(val int) {
}

func (s *Stack) Pop() (int, bool) {
	return 0, false
}

func (s *Stack) Peek() (int, bool) {
	return 0, false
}

func (s *Stack) IsEmpty() bool {
	return false
}
