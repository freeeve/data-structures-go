package queue

type queue struct {
	array []int
}

func NewQueue() *queue {
	q := queue{}
	q.array = []int{}
	return &q
}

func (q *queue) Enqueue(val int) {
	q.array = append(q.array, val)
}

func (q *queue) Dequeue() (int, bool) {
	if len(q.array) == 0 {
		return 0, false
	}
	val := q.array[0]
	q.array = q.array[1:]
	return val, true
}

func (q *queue) Peek() (int, bool) {
	if len(q.array) == 0 {
		return 0, false
	}
	val := q.array[0]
	return val, true
}

func (q *queue) IsEmpty() bool {
	if len(q.array) == 0 {
		return true
	}
	return false
}
