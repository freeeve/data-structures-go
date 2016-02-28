package queue

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type QueueSuite struct{}

var _ = Suite(&QueueSuite{})

func (s *QueueSuite) TestInterface(c *C) {
	var _ Queue = NewQueue()
}

func (s *QueueSuite) TestEnqueue(c *C) {
	queue := NewQueue()
	queue.Enqueue(123)
}

func (s *QueueSuite) TestDequeue(c *C) {
	queue := NewQueue()
	queue.Enqueue(123)
	val, ok := queue.Dequeue()
	c.Assert(val, Equals, 123)
	c.Assert(ok, Equals, true)
}

func (s *QueueSuite) TestDequeueEmpty(c *C) {
	queue := NewQueue()
	val, ok := queue.Dequeue()
	c.Assert(val, Equals, 0)
	c.Assert(ok, Equals, false)
}

func (s *QueueSuite) TestPeek(c *C) {
	queue := NewQueue()
	queue.Enqueue(123)
	val, ok := queue.Peek()
	c.Assert(val, Equals, 123)
	c.Assert(ok, Equals, true)
}

func (s *QueueSuite) TestPeekEmpty(c *C) {
	queue := NewQueue()
	val, ok := queue.Peek()
	c.Assert(val, Equals, 0)
	c.Assert(ok, Equals, false)
}

func (s *QueueSuite) TestIsEmpty(c *C) {
	queue := NewQueue()
	empty := queue.IsEmpty()
	c.Assert(empty, Equals, true)
}

func (s *QueueSuite) TestIsEmptyFalse(c *C) {
	queue := NewQueue()
	queue.Enqueue(123)
	empty := queue.IsEmpty()
	c.Assert(empty, Equals, false)
}

func (s *QueueSuite) TestLeak(c *C) {

}
