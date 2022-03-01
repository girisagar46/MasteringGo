package main

import (
	"fmt"
	"sort"
)

type byLength []string // byLength type is just an alias for the builtin []string type.

// Len, Swap, and Less are from Sort interface.

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "pineapple", "apple"}
	sort.Sort(byLength(fruits)) // convert the original fruit slice to a byLength slice.
	fmt.Println(fruits)
}

/* Output:
[peach apple pineapple]
*/
