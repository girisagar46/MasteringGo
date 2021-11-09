package main

import "fmt"

// More about slice here: https://go.dev/blog/slices-intro

func main() {
	// Different ways to declare slice
	s1 := []int{3, 4, 5} // initialize with values
	s2 := []int{} // initialize but with empty values. Memory is allocated for s2
	var s3 []int // This is preferred way. Declare a slice but do not allocate memory just yet.
	s4 := make([]int, 5, 10) // initialize with no values, allocate memory, specify length of slice and the capacity of underlying array
	fmt.Println(s1, s2, s3, s4)

	// Copying
	x := []int{1, 2, 3, 4, 5}
	x1 := make([]int, 2)
	n := copy(x1, x[2:4])
	fmt.Println("Num of elements copied: ", n)
	x1[0] = 100 // this won't affect x because x1 is a copy
	fmt.Println(x)
	fmt.Println(x1)

	// dynamically expanding a slice
	slice := []int{1, 2, 3}
	slice1 := append(slice, 4, 5, 6)
	slice3 := []int{7, 8, 9}
	slice4 := append(slice1, slice3...) // appending 2 arrays. Remember 3 dots.
	fmt.Println(slice4) // [1 2 3 4 5 6 7 8 9]

	// Slice tricks
	a := []int{100, 200, 300, 400, 500, 600, 700, 800, 900}
	// remove multiple items
	a = append(a[:3], a[5:]...)
	fmt.Println(a) // [100 200 300 600 700 800 900]

	// remove one item
	b := []int{100, 200, 300, 400, 500, 600, 700, 800, 900}
	b = append(b[:3], b[4:]...)
	fmt.Println(b) // [100 200 300 500 600 700 800 900]

}
