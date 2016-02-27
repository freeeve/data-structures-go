package main

type Queue struct {
}

func NewQueue() Queue {
	return Queue{}
}

func (s *Queue) Enqueue(val int) {
}

func (s *Queue) Dequeue() (int, bool) {
	return 0, false
}

func (s *Queue) Peek() (int, bool) {
	return 0, false
}

func (s *Queue) IsEmpty() bool {
	return false
}
