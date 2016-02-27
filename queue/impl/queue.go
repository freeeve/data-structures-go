package main

type Queue struct {
	array []int
}

func NewQueue() Queue {
	q := Queue{}
	q.array = []int{}
	return s
}

func (s *Queue) Enqueue(val int) {
	q.array = append(q.array, val)
}

func (s *Queue) Dequeue() (int, bool) {
	if len(q.array) == 0 {
		return 0, false
	}
	val := q.array[0]
	q.array = q.array[1:]
	return val, true
}

func (s *Queue) Peek() (int, bool) {
	if len(q.array) == 0 {
		return 0, false
	}
	val := q.array[0]
	return val, true
}

func (s *Queue) IsEmpty() bool {
	if len(q.array) == 0 {
		return true
	}
	return false
}
