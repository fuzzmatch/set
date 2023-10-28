package set

type (
	Set struct {
		hash map[string]nothing
	}

	nothing struct{}
)

// Create a new set
func New(initial ...string) *Set {
	s := &Set{make(map[string]nothing)}

	for _, v := range initial {
		s.Insert(v)
	}

	return s
}

// Find the difference between two sets
func (this *Set) Difference(set *Set) *Set {
	n := make(map[string]nothing)

	for k := range this.hash {
		if _, exists := set.hash[k]; !exists {
			n[k] = nothing{}
		}
	}

	return &Set{n}
}

// Call f for each item in the set
func (this *Set) Do(f func(string)) {
	for k, _ := range this.hash {
		f(k)
	}
}

// Test to see whether or not the element is in the set
func (this *Set) Has(element string) bool {
	_, exists := this.hash[element]
	return exists
}

// Add an element to the set
func (this *Set) Insert(element string) {
	this.hash[element] = nothing{}
}

// Find the intersection of two sets
func (this *Set) Intersection(set *Set) *Set {
	n := make(map[string]nothing)

	for k, _ := range this.hash {
		if _, exists := set.hash[k]; exists {
			n[k] = nothing{}
		}
	}

	return &Set{n}
}

// Return the number of items in the set
func (this *Set) Len() int {
	return len(this.hash)
}

// Test whether or not this set is a proper subset of "set"
func (this *Set) ProperSubsetOf(set *Set) bool {
	return this.SubsetOf(set) && this.Len() < set.Len()
}

// Remove an element from the set
func (this *Set) Remove(element string) {
	delete(this.hash, element)
}

// Test whether or not this set is a subset of "set"
func (this *Set) SubsetOf(set *Set) bool {
	if this.Len() > set.Len() {
		return false
	}
	for k, _ := range this.hash {
		if _, exists := set.hash[k]; !exists {
			return false
		}
	}
	return true
}

// Find the union of two sets
func (this *Set) Union(set *Set) *Set {
	n := make(map[string]nothing)

	for k := range this.hash {
		n[k] = nothing{}
	}
	for k := range set.hash {
		n[k] = nothing{}
	}

	return &Set{n}
}

// Convert Set to Slice of a specific type
func (set *Set) ToSlice() []string {
	if set.Len() == 0 {
		return []string{}
	}

	slice := make([]string, set.Len())
	i := 0
	for key := range set.hash {
		slice[i] = key
		i++
	}
	return slice
}
