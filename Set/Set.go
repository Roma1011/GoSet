package Set

import "errors"

type StringSet struct {
	Set []string
}

// Union returns a new set that contains all unique elements from both input sets.
// A∪B={x∣x∈A∨x∈B}
func (StringSet) Union(firstSet StringSet, secondSet StringSet) StringSet {
	var setA StringSet
	setA.Set = append(setA.Set, firstSet.Set...)
	for _, elem := range secondSet.Set {
		if !setA.Contains(elem) {
			setA.Set = append(setA.Set, elem)
		}
	}
	return setA
}

// Intersect returns a new set that contains only the elements that are common in both input sets.
// A∩B={x∣x∈A∧x∈B}
func (StringSet) Intersect(firstSet StringSet, secondSet StringSet) StringSet {
	var resultSet StringSet

	for _, elem := range firstSet.Set {
		if secondSet.Contains(elem) {
			resultSet.Set = append(resultSet.Set, elem)
		}
	}

	return resultSet
}

// Empty Tests if a set is empty../*
func (set *StringSet) Empty() bool {
	if len(set.Set) > 0 {
		return false
	}
	return true
}

// Contains Check if there's an element with the specified key in the set./*
// If the set is empty, it will return true.
func (set StringSet) Contains(item string) bool {

	if set.Empty() {
		return true
	}

	for _, v := range set.Set {
		if v == item {
			return true
		}
	}
	return false
}

// Insert Inserts an element or a range of elements into a set./*
// If the element already exists, it will not be added.
func (set *StringSet) Insert(param string, index int) error {
	if index > len(set.Set) {
		return errors.New("index out of range")
	}
	if set.Contains(param) {
		return nil
	}
	set.Set[index] = param
	return nil
}

// Subset checks if all elements in the second set are contained in the first set.
// A⊆B ⇔ ∀x(x∈A→x∈B)
func (set StringSet) Subset(comparing []string) bool {
	for _, elem := range comparing {
		if !set.Contains(elem) {
			return false
		}
	}
	return true
}
