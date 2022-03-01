package main

import "fmt"

func main() {
	var a = "initial" // no need to define type as it's implicitly inferred by go
	fmt.Println(a)

	var b, c int = 1, 2 // multiple variable declaration at once
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int // this is zero-valued initialization. zero value for int type is 0
	fmt.Println(e)

	// := syntax is shorthand for declaring and initializing a variable
	f := "apple" // implicit type declaration
	fmt.Println(f)
}

/**
initial
1 2
true
0
apple
*/
