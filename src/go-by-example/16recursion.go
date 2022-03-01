package main

import "fmt"

func main() {
	fmt.Println("factorial of 7: ", factorial(7))

	// closures as recursive function
	/*
		Closures can also be recursive, but this requires the closure to be declared with a typed var explicitly before it’s defined.
	*/
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2) // go is smart enough to know which fib to call here
	}
	fmt.Println("7th fibonacci is: ", fib(7))
}

func factorial(n int) int {
	if n == 0 { // in recursive function do not forget to return on base condition
		return 1
	}
	return n * factorial(n-1) // recursive call to factorial function itself
}
