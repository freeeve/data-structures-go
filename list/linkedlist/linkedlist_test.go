package main

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type StackSuite struct{}

var _ = Suite(&StackSuite{})

func (s *StackSuite) TestPush(c *C) {
	stack := NewStack()
	stack.Push(123)
}

func (s *StackSuite) TestPop(c *C) {
	stack := NewStack()
	stack.Push(123)
	val, ok := stack.Pop()
	c.Assert(val, Equals, 123)
	c.Assert(ok, Equals, true)
}

func (s *StackSuite) TestPopEmpty(c *C) {
	stack := NewStack()
	val, ok := stack.Pop()
	c.Assert(val, Equals, 0)
	c.Assert(ok, Equals, false)
}

func (s *StackSuite) TestPeek(c *C) {
	stack := NewStack()
	stack.Push(123)
	val, ok := stack.Peek()
	c.Assert(val, Equals, 123)
	c.Assert(ok, Equals, true)
}

func (s *StackSuite) TestPeekEmpty(c *C) {
	stack := NewStack()
	val, ok := stack.Peek()
	c.Assert(val, Equals, 0)
	c.Assert(ok, Equals, false)
}

func (s *StackSuite) TestLeak(c *C) {

}
