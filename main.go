package main

import (
	"Set/Set"
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var setA Set.StringSet = Set.StringSet{}
	var setB Set.StringSet = Set.StringSet{}
	//-------------------------------------------------------
	fmt.Print("Warning: Please use 'space' for object separator")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Set A Objects: ")
	text, _ := reader.ReadString('\n')
	var buildString []string = SetBuilder(text)
	setA.Set = make([]string, (len(buildString)*2)-(DelimiterCount(32, text))-1)
	for i, str := range buildString {
		setA.Insert(str, i)
	}
	sort.Strings(setA.Set)
	//-------------------------------------------------------
	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Enter Set B Objects: ")
	text, _ = reader.ReadString('\n')
	buildString = SetBuilder(text)
	setB.Set = make([]string, (len(buildString)*2)-(DelimiterCount(32, text))-1)
	for i, str := range buildString {
		setB.Insert(str, i)
	}
	sort.Strings(setB.Set)
	setA.Power()
	//--------------------------------------------------------
	var isSubset bool = setA.Subset(setB.Set)
	fmt.Println("Is Set A a subset of Set B? ", isSubset)
}
