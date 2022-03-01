package main

import "fmt"

/**
There's only one looping mechanism in go. and that's for loop.
*/

func main() {

	fmt.Println("some basis for loops type:")

	i := 1
	for i <= 3 { // most basic form with the single condition. compares to while loop in other languages.
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 5; j++ { // the classic init/condition/increment style
		fmt.Println(j)
	}

	for { // if for loop does not have any conditions, then it's considered as infinite loop.
		fmt.Println("infinite loop")
		break // make sure to break the infinite loop. If the break is not present here, then this loop will run forever
	}

	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue // continue will continue the for loop to next iteration of loop
		}
		fmt.Println(n)
	}
}

/**
some basis for loops type:
1
2
3
0
1
2
3
4
infinite loop
1
3
*/
