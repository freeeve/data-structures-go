package hashset

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type HashSetSuite struct{}

var _ = Suite(&HashSetSuite{})

func (s *HashSetSuite) TestInterface(c *C) {
	var _ HashSet = NewHashSet()
}

func (s *HashSetSuite) TestAdd(c *C) {
	hashset := NewHashSet()
	ok := hashset.Add(123)
	c.Assert(ok, Equals, true)
}

func (s *HashSetSuite) TestAddFail(c *C) {
	hashset := NewHashSet()
	ok := hashset.Add(123)
	ok = hashset.Add(123)
	c.Assert(ok, Equals, false)
}

func (s *HashSetSuite) TestContains(c *C) {
	hashset := NewHashSet()
	hashset.Add(123)
	contains := hashset.Contains(123)
	c.Assert(contains, Equals, true)
}

func (s *HashSetSuite) TestContainsFalse(c *C) {
	hashset := NewHashSet()
	contains := hashset.Contains(123)
	c.Assert(contains, Equals, false)
}

func (s *HashSetSuite) TestIsEmpty(c *C) {
	hashset := NewHashSet()
	empty := hashset.IsEmpty()
	c.Assert(empty, Equals, true)
}

func (s *HashSetSuite) TestIsEmptyFalse(c *C) {
	hashset := NewHashSet()
	hashset.Add(123)
	empty := hashset.IsEmpty()
	c.Assert(empty, Equals, false)
}

func (s *HashSetSuite) TestLeak(c *C) {

}
