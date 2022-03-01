package main

import "fmt"

/**
NOTE: There is no ternary if in Go. So we need to use full if statement even for small basic conditions
*/

func main() {

	// if statement without an else
	if 8%2 == 0 {
		fmt.Println("8 is divisible by 2")
	}

	for i := 0; i < 5; i++ {
		// if statement with else
		if i%2 == 0 {
			fmt.Printf("%v is even\n", i)
		} else {
			fmt.Printf("%v is odd\n", i)
		}
	}

	// we can define a variable in the if-statement
	if num := 9; num < 0 { // here num variable is available to all branches
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// the scope of `num` is only inside the if statements
	//fmt.Println(num) // throws error undefined: num
}

/**
8 is divisible by 2
0 is even
1 is odd
2 is even
3 is odd
4 is even
9 has 1 digit
*/
