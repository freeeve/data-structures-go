package queue

type queue struct {
}

func NewQueue() *queue {
	return nil
}

func (q *queue) Enqueue(val int) {
}

func (q *queue) Dequeue() (int, bool) {
	return 0, false
}

func (q *queue) Peek() (int, bool) {
	return 0, false
}

func (q *queue) IsEmpty() bool {
	return false
}
