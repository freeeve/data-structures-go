package queue

type Queue interface {
	Enqueue(val int)
	Dequeue() (int, bool)
	Peek() (int, bool)
	IsEmpty() bool
}
