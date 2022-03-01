package main

import (
	"fmt"
	"sort"
)

func main() {
	// Note that sorting is in-place, so it changes the given slice and does not return a new one.
	// Sort methods are specific to the builtin type; here’s an example for strings.
	strs := []string{"peach", "apple", "pear", "plum"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4, 5, 1, 3, 9}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	s := sort.IntsAreSorted(ints) // We can also use sort to check if a slice is already in sorted order.
	fmt.Println("Sorted: ", s)
	fmt.Println("is sorted? ", sort.IntsAreSorted([]int{1, 2, 3, 4, 10, 6, 7})) // false
}
