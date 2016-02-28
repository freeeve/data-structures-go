package hashset

type HashSet interface {
	Add(val int) bool
	Contains(val int) bool
	Remove(val int) bool
	IsEmpty() bool
}

type hashSet struct {
}

func NewHashSet() *hashSet {
	return nil
}

func (s *hashSet) Add(val int) bool {
	return false
}

func (s *hashSet) Contains(val int) bool {
	return false
}

func (s *hashSet) Remove(val int) bool {
	return false
}

func (s *hashSet) IsEmpty() bool {
	return false
}
