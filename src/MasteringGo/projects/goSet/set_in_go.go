package main

import "fmt"

type Set map[string]struct{} // map that maps to empty struct

// IDEA: Map keys are unique, and we can leverage this to achieve set behaviour
// We define empty struct `struct{}` as map value because empty struct do not consume any memory
func main() {
	s := make(Set)
	s["item1"] = struct{}{}
	s["item2"] = struct{}{}
	s["item2"] = struct{}{}
	fmt.Println(getSetValues(s)) // Just outputs: [item2 item1] even though there are 2 item2
}

func getSetValues(s Set) []string {
	var retVal []string
	for k, _ := range s {
		retVal = append(retVal, k)
	}
	return retVal
}
