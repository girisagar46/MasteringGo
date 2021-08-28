package utility

import "fmt"

func Looping() {
	n := 1
	for i := 10; i > 0; i-- {
		n *= i
	}
	fmt.Println("Result: ", n)

	// There's no while loop in Go
	n = 1
	for n <= 50 {
		n += n
	}
	fmt.Println("Result: ", n)

	// Infinite loop
	// WARN: DANGER
	/*
	   for {
	       fmt.Println("Forever and ever")
	   }
	*/

	// defer inside loop
	for i := 0; i < 5; i++ {
		// WARN: Do not do defer inside loop
		// ERROR: Possible resource leak, 'defer' is called in the 'for' loop
		defer fmt.Print(i, " ") // Executed with LIFO basis. Output: 4 3 2 1 0
	}
	fmt.Println("Done")
}
