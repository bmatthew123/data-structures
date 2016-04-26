package ds

import (
	"testing"
)

type TestStruct struct {
	number  int
	boolean bool
	str     string
	float   float64
}

func TestAddFloat(t *testing.T) {
	s := NewSet()
	s.Add(153.82)
	if _, ok := s.items[153.82]; !ok {
		t.Fatal("TestAddFloat failed. Set should contain 153.82")
	}
}

func TestAddStruct(t *testing.T) {
	s := NewSet()
	s.Add(TestStruct{
		number:  1234,
		boolean: false,
		str:     "Some string value",
		float:   74.1234,
	})
	if _, ok := s.items[TestStruct{1234, false, "Some string value", 74.1234}]; !ok {
		t.Fatal("TestAddStruct failed. Set should contain the struct.")
	}
}

func TestAddDifferent(t *testing.T) {
	s := NewSet()
	s.Add(153.82)
	s.Add(TestStruct{
		number:  1234,
		boolean: false,
		str:     "Some string value",
		float:   74.1234,
	})
	if _, ok := s.items[153.82]; !ok {
		t.Fatal("TestAddDifferent failed. Set should contain 153.82")
	}
	if _, ok := s.items[TestStruct{1234, false, "Some string value", 74.1234}]; !ok {
		t.Fatal("TestAddDifferent failed. Set should contain the struct.")
	}
}

func TestCountZero(t *testing.T) {
	s := NewSet()
	if s.Count() != 0 {
		t.Fatal("Set should initially have 0 count")
	}
}

func TestCount(t *testing.T) {
	s := NewSet()
	for i := 0; i < 5; i++ {
		s.Add(i)
	}
	if s.Count() != 5 {
		t.Fatal("Set should initially have count 5")
	}
}

func TestContains(t *testing.T) {
	s := NewSet()
	s.Add(1)
	s.Add(TestStruct{
		number:  1234,
		boolean: false,
		str:     "Some string value",
		float:   74.1234,
	})
	if !s.Contains(1) {
		t.Fatal("Set should contain integer 1")
	}
	ts := TestStruct{number: 1234,
		boolean: false,
		str:     "Some string value",
		float:   74.1234,
	}
	if !s.Contains(ts) {
		t.Fatal("Set should contain the struct")
	}
}

func TestRemove(t *testing.T) {
	s := NewSet()
	s.Remove(1) // Shouldn't throw an error
	s.Add(1)
	s.Remove(1)
	if s.Count() != 0 {
		t.Fatal("Set should be empty")
	}
}

func TestContents(t *testing.T) {
	s := NewSet()
	s.Add(1)
	s.Add(TestStruct{
		number:  1234,
		boolean: false,
		str:     "Some string value",
		float:   74.1234,
	})
	contents := s.Contents()
	if !(contents[0] == 1 || contents[1] == 1) {
		t.Fatal("Contents should include a 1")
	}
	ts := TestStruct{number: 1234,
		boolean: false,
		str:     "Some string value",
		float:   74.1234,
	}
	if !(contents[0] == ts || contents[1] == ts) {
		t.Fatal("Contents should include the struct")
	}
}
