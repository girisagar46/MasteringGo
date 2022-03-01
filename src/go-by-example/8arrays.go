package main

import "fmt"

func main() {
	var a [5]int // create empty array that holds exactly 5 elements. These elements are 0 valued.
	fmt.Println("empty array: ", a)
	fmt.Println("len : ", len(a))      // built-in len function returns length of array
	fmt.Println("capacity : ", cap(a)) // built-in cap function returns the capacity of the array

	a[4] = 100 // using array index to set the value
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])
	fmt.Println("length: ", len(a))

	b := [5]int{1, 2, 3, 4, 5} // declare array and initialize with values
	//b := [5]int{1, 2, 3, 4, "5"} // BAD: we can not mix match type in array. this declared as int storage
	fmt.Println("dcl: ", b)

	var twoDimensionalArray [2][3]int // all values initialized to zero
	fmt.Println("2D array before: ", twoDimensionalArray)
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDimensionalArray[i][j] = i + j
		}
	}
	fmt.Println("2D array after: ", twoDimensionalArray)
}

/**
empty array:  [0 0 0 0 0]
len :  5
capacity :  5
set:  [0 0 0 0 100]
get:  100
length:  5
dcl:  [1 2 3 4 5]
2D array before:  [[0 0 0] [0 0 0]]
2D array after:  [[0 1 2] [1 2 3]]
*/
