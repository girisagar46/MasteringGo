package main

import "fmt"

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroVal(i)
	fmt.Println("zeroVal:", i)

	zeroPtr(&i)
	fmt.Println("zeroPtr:", i)

	// zeroVal does not change the i in main,
	// but zeroPtr does because it has a reference to the memory address for that variable.

	fmt.Println("pointer:", &i)
}

func zeroVal(iVal int) { // parameter type is int, so it'll be pass by value. zeroVal will receive a copy
	fmt.Println("address of iVal:", &iVal)
	iVal = 0
}

func zeroPtr(iPtr *int) { // parameter type is *int, so it'll be pass by reference. zeroPtr will receive address of i
	*iPtr = 0 // this mutation will also change the value of i in main function
}

/*
initial: 1
address of iVal: 0x140000160c8
zeroVal: 1
zeroPtr: 0
pointer: 0x140000160b8
*/
