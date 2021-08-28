package main

import (
	"MasteringGo/src/myfirstapp/utility"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	utility.SayHi()
	//utility.sayHello() // this function is not accessible here

	res, computeRes := utility.Inner()
	fmt.Println(res, "", computeRes) // 3  50

	inc := utility.Increment()
	fmt.Println(inc(), inc(), inc()) // 1 2 3

	// Variables
	var v1 int   // variable declaration
	v1 = 100     // assign value to variable v1
	var v2 = 200 // declare and initialize variable (Should use var if declaring variable outside the function)
	v3 := 300    // declare and initialize variable with type inference (This declaration only works inside the function)
	fmt.Println(v1, v2, v3)

	// Arrays are not flexible in Go, as we need to define size beforehand
	var a [2]int
	a[0], a[1] = 4, 5
	fmt.Println("Array a: ", a, len(a))

	b := [3]int{90, 99} // Empty space will be filled with zero because it's of int type
	fmt.Println("Array b: ", b, len(b))

	// Slice is more flexible than Arrays
	c := []int{99, 55} // If we do not define size of Array, it becomes a slice
	c = append(c, 88)  // append takes slice as first input
	c = append(c, 45)
	fmt.Println("Slice c: ", c, len(c))

	// Pointers
	var p1 *int
	num := 4
	p1 = &num
	fmt.Println(p1) // p1 stores tha address of num

	p2 := &num
	fmt.Println(p2)  // prints address where num is stored
	fmt.Println(*p2) // prints the value of num. This is Dereferenceing the pointer value to the original one

	myVal := 5
	fmt.Println("myVal Before", myVal)
	utility.ChangeMe(&myVal) // We pass the address of myVal and in the ChangeMe function, we're deferencing it
	fmt.Println("myVal After", myVal)

	// type casting
	n := 32
	f := float64(n)
	const pi = 3.1415 // This is constant value
	fmt.Println(pi * f)
	// Enum constants
	const (
		big   int = 1 << 10
		small     = big >> 99
	)
	fmt.Println(big)
	// big = 100 // this throws error because big is the constant.
	fmt.Println(small)

	// Flow Control in go. See utility.flowControl.go
	utility.FlowControl()

	// Loops in Go. See utility.loops.go
	utility.Looping()
}
