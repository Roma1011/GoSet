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

// Empty Tests if a set is empty../*
func (set *Set) Empty() bool {
	if len(set.set) > 0 {
		return false
	}
	return true
}

// Contains Check if there's an element with the specified key in the set./*
func (set *Set) Contains(item string) bool {

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

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Set A Objects: ")
	text, _ := reader.ReadString('\n')
	var buildString []string = SetBuilder(text)

	var n int = (len(buildString) * 2) - (DelimiterCount(32, text)) - 1
	println(n)
	setA.set = make([]string, (len(buildString)*2)-(DelimiterCount(32, text))-1)

	for i, str := range buildString {
		setA.Insert(str, i)
	}
	sort.Strings(setA.set)
	println(setA.Contains("A"))
}
