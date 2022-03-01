package main

import "fmt"

func main() {
	a := 100
	b := 200
	sum := add(a, b)
	fmt.Printf("%v + %v = %v\n", a, b, sum)

	product := multiply(a, b)
	fmt.Printf("%v * %v = %v\n", a, b, product)
}

func multiply(a, b int) int { // if the parameters are of same type, we can only put type at the end
	return a * b
}

func add(a int, b int) int { // Go requires explicit return
	return a + b
}

/**
100 + 200 = 300
100 * 200 = 20000
*/
