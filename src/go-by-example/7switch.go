package main

import (
	"fmt"
	"time"
)

func main() {

	// the basic switch statement
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // multiple expressions in the same case statement.
		fmt.Println("It's the Weekend! 🎉") // unicode and emoji are supported in go
	default:
		fmt.Println("It's the Weekday! 😭")
	}

	t := time.Now()
	switch { // switch without an expression is an alternate way to express if/else logic
	case t.Hour() < 12: // case can also have a condition expression. It's only possible for empty switch
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) { // we can also have variable assignment in switch statement.
		case bool:
			fmt.Println("I am of bool type")
		case int:
			fmt.Println("I am of int type")
		default:
			fmt.Printf("I don't know who am I! %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(0)
	whatAmI("Hello")
}

//Note: Some output depends on when you run it
/**
Write 2 as Two
It's the Weekend! 🎉
It's after noon
I am of bool type
I am of int type
I don't know who am I! string
*/
