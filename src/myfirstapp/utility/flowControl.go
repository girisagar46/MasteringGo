package utility

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func FlowControl() {


	// if statement

	// Here the value of i is visible to all blocks below if
	// Also, the variable i is local to this statement
	if i := randomNum(); i < 0 {
		fmt.Println("i is a negative number. i.e. ", i)
	} else if i == 0 {
		fmt.Println("i is Zero i.e. ", i)
	} else {
		fmt.Println("i is positive number i.e. ", i)
	}

	// switch statement
	// There's no break in each case. break is implicit
	switch i := randomNum(); {
	case i < 0:
		fmt.Println("i is a negative number. i.e. ", i)
	case i == 0:
		fmt.Println("i is Zero i.e. ", i)
	default:
		fmt.Println("i is positive number i.e. ", i)
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
		fallthrough // this will remove the effect of case just below it
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
		// if default condition matches, this will execute at the end
		defer fmt.Println("Exiting FlowControl function...")
		fmt.Println("Entering function")
	}

}


func randomNum() (randNum int) {
	rand.Seed(time.Now().UnixNano())
	min := -5
	max := 5
	randNum = rand.Intn(max - min + 1) + min
	return
}
