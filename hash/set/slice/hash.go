package hashset

type HashSet interface {
	Add(val int) bool
	Contains(val int) bool
	Remove(val int) bool
	IsEmpty() bool
}

type hashSet struct {
	array []*int
	size  int
}

func NewHashSet() *hashSet {
	h := hashSet{}
	h.array = []*int{}
	return &h
}

func hashInt(i int) int {
	return i
}

func (s *hashSet) Add(val int) bool {
	if len(s.array) == 0 || float64(s.size)/float64(len(s.array)) > 0.5 {
		s.growArray()
	}
	bucket := hashInt(val) % len(s.array)
	for s.array[bucket] != nil && *s.array[bucket] != val {
		bucket++
		if bucket > len(s.array)-1 {
			bucket = 0
		}
	}
	if s.array[bucket] == nil {
		s.array[bucket] = &val
		s.size++
	} else {
		return false
	}
	return true
}

func (s *hashSet) growArray() {
	if len(s.array) == 0 {
		s.array = make([]*int, 256, 256)
	} else {
		temp := s.array
		s.array = make([]*int, len(s.array)*2)
		s.size = 0
		for _, x := range temp {
			if x != nil {
				s.Add(*x)
			}
		}
	}
}

func (s *hashSet) Contains(val int) bool {
	if len(s.array) == 0 {
		return false
	}
	bucket := hashInt(val) % len(s.array)
	for s.array[bucket] != nil && *s.array[bucket] != val {
		bucket++
		if bucket > len(s.array)-1 {
			bucket = 0
		}
	}
	if s.array[bucket] == nil {
		return false
	} else {
		return true
	}
}

func (s *hashSet) Remove(val int) bool {
	bucket := hashInt(val) % len(s.array)
	for s.array[bucket] != nil && *s.array[bucket] != val {
		bucket++
		if bucket > len(s.array)-1 {
			bucket = 0
		}
	}
	if s.array[bucket] == nil {
		return false
	} else {
		s.array[bucket] = nil
		s.size--
		return true
	}
}

func (s *hashSet) IsEmpty() bool {
	return s.size == 0
}
