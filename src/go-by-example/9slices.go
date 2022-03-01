package main

import "fmt"

/**
Slices are a key data type in Go, giving a more powerful interface to sequence than arrays.
Inside slice: https://chidiwilliams.com/post/inside-a-go-slice/
*/
func main() {
	s := make([]string, 3) // make a slice of strings of length 3 (initially zero-valued).
	fmt.Println("empty slice", s)
	fmt.Printf("type of s is %T\n", s)
	fmt.Println("len: ", len(s))
	fmt.Println("cap: ", cap(s))
	// setting values to slice
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2]) // getting values

	//s[4] = "Hello" // panic: runtime error: index out of range [4] with length
	fmt.Println("size of s: ", len(s)) // len returns the length of slice as expected

	// append is built-in support for slice to append
	// elements are appended at the end of slice
	// append returns a slice containing one or more new values.
	// we need to accept the new slice value as append returns a slice
	// append can be used to dynamically increase the size of the data structure
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("size of s after append: ", len(s))
	fmt.Println("capacity of s after append: ", cap(s))
	fmt.Println(s)

	// Slices can also be copy’d. Here we create an empty slice c of the same length as s and copy into c from s.
	c := make([]string, len(s))
	copy(c, s) // signature is: copy(dst []Type, src []Type)
	fmt.Println("copy: ", c)
	fmt.Println("size of c: ", len(s))
	fmt.Println("capacity of c: ", cap(s))

	// slice supports slice operator (similar to python)
	l := c[2:5] // c[low_index:high_index]. The high_index is exclusive. l contains s[2], s[3], and s[4]
	fmt.Println("slice1 c[2:5]", l)

	l = c[:5] // here low_index is 0 and high index is 4. l contains s[0], s[1], s[2], s[3] and s[4]
	fmt.Println("slice2 c[:5]", l)

	l = c[2:] // here low_index is 0 and high index is len(c)-1. l contains s[2] ... s[5]
	fmt.Println("slice3 c[2:]", l)

	// We can declare and initialize a variable for slice in a single line as well.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	u := []string{"j", "k", "l"}
	// we can merge two slices, but we need to unpack the 2nd slice
	merged := append(t, u...)
	fmt.Println("Merging two slices", merged)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	fmt.Println("Difference between an Array and Slice")
	// Read more here: https://go.dev/blog/slices-intro
	arr := [5]int{1, 2, 3, 4} // An array type definition specifies a length and an element type.
	fmt.Println(arr)          // [1 2 3 4 0]
	//fmt.Println(append(arr, 5)) // not possible because append does not work with array

	sli := []int{1, 2, 3, 4}    // this is a slice literal. Unlike an array type, a slice type has no specified length.
	fmt.Println(sli)            // [1 2 3 4]
	fmt.Println(append(sli, 5)) // perfectly okay with slice

	sli2 := make([]int, 4) // we can use the built-in make function to create a slice. here length = capacity
	fmt.Println("Sli2", "len: ", len(sli2), " cap: ", cap(sli2), " elements: ", sli2)

	sli3 := make([]int, 1, 2) // here length = 1, and capacity = 2
	fmt.Println("Sli3", "len: ", len(sli3), " cap: ", cap(sli3), " elements: ", sli3)

	x := [3]string{"Лайка", "Белка", "Стрелка"}
	// converting an array to slice
	xSlice := x[:] // a slice referencing the storage of x
	xSlice = append(xSlice, "Sagar")
	fmt.Println(xSlice)

}

/*
empty slice [  ]
type of s is []string
len:  3
cap:  3
set: [a b c]
get: c
size of s:  3
size of s after append:  6
capacity of s after append:  6
[a b c d e f]
copy:  [a b c d e f]
size of c:  6
capacity of c:  6
slice1 c[2:5] [c d e]
slice2 c[:5] [a b c d e]
slice3 c[2:] [c d e f]
dcl: [g h i]
Merging two slices [g h i j k l]
2d:  [[0] [1 2] [2 3 4]]
Difference between an Array and Slice
[1 2 3 4 0]
[1 2 3 4]
[1 2 3 4 5]
Sli2 len:  4  cap:  4  elements:  [0 0 0 0]
Sli3 len:  1  cap:  2  elements:  [0]
[Лайка Белка Стрелка Sagar]
*/
