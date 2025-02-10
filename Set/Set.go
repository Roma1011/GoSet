package Set

import (
	"errors"
	"math"
)

type StringSet struct {
	Set []string
}

// Union returns a new set that contains all unique elements from both input sets.
// A∪B={x∣x∈A∨x∈B}
func (StringSet) Union(firstSet StringSet, secondSet StringSet) StringSet {
	var setA StringSet
	setA.Set = append(setA.Set, firstSet.Set...)
	for _, elem := range secondSet.Set {
		if !setA.Belongs(elem) {
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
		if secondSet.Belongs(elem) {
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

// Belongs Check if there's an element with the specified key in the set./*
// x ∈ B
func (set StringSet) Belongs(item string) bool {

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
	if set.Belongs(param) {
		return nil
	}
	set.Set[index] = param
	return nil
}

// Subset checks if all elements in the second set are contained in the first set.
// A⊆B ⇔ ∀x(x∈A→x∈B)
func (set StringSet) Subset(comparing []string) bool {
	for _, elem := range comparing {
		if !set.Belongs(elem) {
			return false
		}
	}
	return true
}

// Power P(A)
// all subsets of A
func (set StringSet) Power() {
	var iterator float64 = math.Pow(2, float64(len(set.Set)))
	for i := 0; i < int(iterator); i++ {
		for j := 0; j < len(set.Set); j++ {
			if i&(1<<uint(j)) > 0 {
				print(set.Set[j], " ")
			}
		}
		println()
	}
}
