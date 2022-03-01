package main

import "fmt"

func main() {
	sum, product := sumProduct(100, 200)
	fmt.Println("sum = ", sum)
	fmt.Println("product = ", product)
}

func sumProduct(a, b int) (int, int) { // The (int, int) in this function signature shows that the function returns 2 ints.
	return a + b, a * b
}
