package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Set struct {
	set []string
}

// Union returns a new set that contains all unique elements from both input sets.
// A∪B={x∣x∈A∨x∈B}
func (Set) Union(firstSet Set, secondSet Set) Set {
	var setA Set
	setA.set = append(setA.set, firstSet.set...)
	for _, elem := range secondSet.set {
		if !setA.Contains(elem) {
			setA.set = append(setA.set, elem)
		}
	}
	return setA
}

// Intersect returns a new set that contains only the elements that are common in both input sets.
// A∩B={x∣x∈A∧x∈B}
func (Set) Intersect(firstSet Set, secondSet Set) Set {
	var resultSet Set

	for _, elem := range firstSet.set {
		if secondSet.Contains(elem) {
			resultSet.set = append(resultSet.set, elem)
		}
	}

	return resultSet
}

// AllContains checks if all elements in the second set are contained in the first set.
// It returns true if all elements of the second set are found in the first set, otherwise false.
func AllContains(firstSet, secondSet []string) bool {
	if len(secondSet) > len(firstSet) {
		firstSet, secondSet = secondSet, firstSet
	}

	for i := 0; i < len(secondSet); i++ {
		found := false
		for j := 0; j < len(firstSet); j++ {
			if secondSet[i] == firstSet[j] {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Empty Tests if a set is empty../*
func (set *Set) Empty() bool {
	if len(set.set) > 0 {
		return false
	}
	return true
}

// Contains Check if there's an element with the specified key in the set./*
func (set Set) Contains(item string) bool {

	if set.Empty() {
		return true
	}

	for _, v := range set.set {
		if v == item {
			return true
		}
	}
	return false
}

// Insert Inserts an element or a range of elements into a set./*
func (set *Set) Insert(param string, index int) error {
	if index > len(set.set) {
		return errors.New("index out of range")
	}
	if set.Contains(param) {
		return nil
	}
	set.set[index] = param
	return nil
}

// SetBuilder takes a space-separated string and splits it into a slice of strings.
// It removes any trailing newline character from the last element in the slice.
func SetBuilder(stdn string) []string {
	var buildIt []string = strings.Split(stdn, " ")
	buildIt[len(buildIt)-1] = strings.ReplaceAll(buildIt[len(buildIt)-1], "\n", "")
	return buildIt
}

// DelimiterCount Function to count occurrences of a specific character
func DelimiterCount(char rune, source string) int {
	var counter int = 0
	for _, value := range source {
		if value == char {
			counter++
		}
	}
	return counter
}

func main() {
	var setA Set = Set{}
	var setB Set = Set{}
	//-------------------------------------------------------
	fmt.Print("Warning: Please use 'space' for object separator")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Set A Objects: ")
	text, _ := reader.ReadString('\n')
	var buildString []string = SetBuilder(text)
	setA.set = make([]string, (len(buildString)*2)-(DelimiterCount(32, text))-1)
	for i, str := range buildString {
		setA.Insert(str, i)
	}
	sort.Strings(setA.set)
	//-------------------------------------------------------
	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Enter Set B Objects: ")
	text, _ = reader.ReadString('\n')
	buildString = SetBuilder(text)
	setB.set = make([]string, (len(buildString)*2)-(DelimiterCount(32, text))-1)
	for i, str := range buildString {
		setB.Insert(str, i)
	}
	sort.Strings(setB.set)
	//--------------------------------------------------------
	_ := setA.Union(setB, setA)
}
