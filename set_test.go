package set

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	s := New[string]()

	s.Insert("5")

	if s.Len() != 1 {
		t.Errorf("Length should be 1")
	}
}
func TestHas(t *testing.T) {
	s := New[string]("5")

	if !s.Has("5") {
		t.Errorf("Membership test failed")
	}
}

func TestRemove(t *testing.T) {
	s := New[string]("5")
	s.Remove("5")

	if s.Len() != 0 {
		t.Errorf("Length should be 0")
	}

	if s.Has("5") {
		t.Errorf("The set should be empty")
	}
}

func TestIntersection(t *testing.T) {
	// Intersection
	s1 := New[string]("1", "2", "3", "4", "5", "6")
	s2 := New[string]("4", "5", "6")
	s3 := s1.Intersection(s2)
	if s3.Len() != 3 {
		t.Errorf("Length should be 3 after intersection")
	}

	if !(s3.Has("4") && s3.Has("5") && s3.Has("6")) {
		t.Errorf("Set should contain only 4, 5, 6")
	}
}

// Difference
func TestDifference(t *testing.T) {
	s1 := New[string]("1", "2", "3", "4", "5", "6")
	s2 := New[string]("4", "5", "6")
	s3 := s1.Difference(s2)
	if s3.Len() != 3 {
		t.Errorf("Length should be 3")
	}

	if !(s3.Has("1") && s3.Has("2") && s3.Has("3")) {
		t.Errorf("Set should only contain 1, 2, 3")
	}
}

// Union
func TestUnion(t *testing.T) {
	s1 := New[string]("7", "8", "9")
	s2 := New[string]("4", "5", "6")
	s3 := s2.Union(s1)

	if s3.Len() != 6 {
		t.Errorf("Length should be 6 after union")
	}

	if !(s3.Has("7")) {
		t.Errorf("Set should contain 4, 5, 6, 7, 8, 9")
	}
}

// Subset
func TestSubsetOf(t *testing.T) {
	s1 := New[string]("1", "2", "3", "4", "5", "6")
	s2 := New[string]("4", "5", "6")
	if !s1.SubsetOf(s1) {
		t.Errorf("set should be a subset of itself")
	}

	if !s2.SubsetOf(s1) {
		t.Errorf("[4,5,6] should be a subset of [1,2,3,4,5,6]")
	}

	if s1.SubsetOf(s2) {
		t.Errorf("[1,2,3,4,5,6] should NOT be a subset of [4,5,6]")
	}
}

// Proper Subset
func TestProperSubsetOf(t *testing.T) {
	s1 := New[string]("1", "2", "3", "4", "5", "6")
	s2 := New[string]("4", "5", "6")

	if s1.ProperSubsetOf(s1) {
		t.Errorf("Set should not be a subset of itself")
	}

	if s1.ProperSubsetOf(s2) {
		t.Errorf("[1,2,3,4,5,6] should not be a subset of [4,5,6]")
	}

	if !s2.ProperSubsetOf(s1) {
		t.Errorf("[4,5,6] should be a subset of [1,2,3,4,5,6]")
	}
}

// Elements
func TestElements(t *testing.T) {
	s1 := New[string]("1")
	slice := s1.Elements()
	expected := []string{"1"}

	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("Failed Converting Set to Slice of Elements.")
	}
}

// Copy
func TestCopy(t *testing.T) {
	s1 := New[string]("1", "2")
	copy := s1.Copy()

	if !s1.Equals(copy) {
		t.Errorf("The copy be equal to the original.")
	}

	s1.Remove("1")

	if !copy.Has("1") {
		t.Errorf("Removing an element from the original should not remove an element from the copy.")
	}

	copy.Remove("2")
	if !s1.Has("2") {
		t.Errorf("Removing an element from the copy should not remove an element from the original.")
	}
}

// Equals
func TestEquals(t *testing.T) {
	s1 := New[string]("1", "2", "3", "4", "5", "6")
	s2 := New[string]("4", "5", "6")
	s3 := New[string]("5", "4", "6")

	if s1.Equals(s2) {
		t.Errorf("[1,2,3,4,5,6] should not be equal to [4,5,6]")
	}

	if !s1.Equals(s1) {
		t.Errorf("A set should be equal to itself")
	}

	if !s2.Equals(s3) {
		t.Errorf("Sets defined separately, but containing the same elements should be equal.")
	}

	if s2.Equals(s3) && !s3.Equals(s2) {
		t.Errorf("The equality operation should be symmetric (a=b implies b=a)")
	}
}
