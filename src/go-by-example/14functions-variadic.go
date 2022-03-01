package main

import "fmt"

// NOTES:
// Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.
func main() {
	fmt.Println(sumMultiple()) // we can call Variadic function with zero argument. similar to fmt.Println()
	fmt.Println(sumMultiple(55))
	fmt.Println(sumMultiple(100, 200))
	fmt.Println(sumMultiple(1, 2, 3, 4, 5))
	nums := []int{1, 2, 3, 4}
	fmt.Println(sumMultiple(nums...))

	//arr := [5]int{1, 2, 3, 4}
	//fmt.Println(sumMultiple(arr...)) // the variadic function only accepts slice

	arr := [5]int{1, 2, 3, 4}
	// if we really want to pass array, we convert array to slice first
	// but be very careful if you're doing multiplication or division because the array's last element is 0
	fmt.Println(sumMultiple(arr[:]...))

}

func sumMultiple(nums ...int) int { // this function only accepts a slice of type int
	fmt.Println("nums : ", nums) // here nums is a slice. (NOT ARRAY)
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

/*
nums :  []
0
nums :  [55]
55
nums :  [100 200]
300
nums :  [1 2 3 4 5]
15
nums :  [1 2 3 4]
10
nums :  [1 2 3 4 0] // this is an array with size 5, but we only initialize upto 4, hence last element is 0
10
*/
