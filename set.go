package ds

type Set struct {
	items map[interface{}]bool
}

func NewSet() (s *Set) {
	s = &Set{}
	s.items = make(map[interface{}]bool, 0)
	return
}

func (s *Set) Add(items ...interface{}) {
	for _, item := range items {
		s.items[item] = true
	}
}

func (s *Set) Remove(items ...interface{}) {
	for _, item := range items {
		delete(s.items, item)
	}
}

func (s *Set) Contains(items ...interface{}) bool {
	for _, item := range items {
		if _, ok := s.items[item]; !ok {
			return false
		}
	}
	return true
}

func (s *Set) Count() int {
	return len(s.items)
}

func (s *Set) Contents() []interface{} {
	c := make([]interface{}, len(s.items))
	counter := 0
	for item, _ := range s.items {
		c[counter] = item
		counter++
	}
	return c
}
