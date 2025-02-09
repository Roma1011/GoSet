package main

import "strings"

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
