package queue

type queue struct {
	head *queueNode
	tail *queueNode
}

type queueNode struct {
	next *queueNode
	prev *queueNode
	val  int
}

func NewQueue() *queue {
	q := queue{}
	return &q
}

func (q *queue) Enqueue(val int) {
	if q.head == nil {
		q.head = &queueNode{prev: nil, next: nil, val: val}
		q.tail = q.head
	} else {
		temp := q.tail
		q.tail = &queueNode{prev: temp, next: nil, val: val}
		temp.next = q.tail
	}
}

func (q *queue) Dequeue() (int, bool) {
	if q.head == nil {
		return 0, false
	}
	temp := q.head
	q.head = temp.next
	q.head.prev = nil
	return temp.val, true
}

func (q *queue) Peek() (int, bool) {
	if q.head == nil {
		return 0, false
	}
	temp := q.head
	return temp.val, true
}

func (q *queue) IsEmpty() bool {
	if q.head == nil {
		return true
	}
	return false
}
