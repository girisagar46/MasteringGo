package main

import (
	"errors"
	"fmt"
)

/*
In Go, it’s idiomatic to communicate errors via an explicit, separate return value. This approach makes it easy to see
which functions return errors and to handle them using the same language constructs employed for any other, non-error tasks.

More on error in Go: https://go.dev/blog/error-handling-and-go
*/

func f1(arg int) (int, error) { // By convention, errors are the last return value and have type error, a built-in interface.
	if arg == 45 {
		return -1, errors.New("can't work with 42") // errors.New constructs a basic error value with the given error message.
	}
	return arg + 3, nil // A nil value in the error position indicates that there was no error.
}

type argError struct { // It’s possible to use custom types as errors by implementing the Error() method on them L25.
	arg     int
	problem string
}

func (e *argError) Error() string { // implement Error method on argError type
	return fmt.Sprintf("%d - %s", e.arg, e.problem)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg: arg, problem: "can't work with it"} // here &argError is of type error
	}
	return arg + 3, nil
}

func main() {
	fmt.Println("Calling f1")
	val, err := f1(34)
	fmt.Println(val, " error: ", err)
	val2, err2 := f1(45)
	fmt.Println(val2, " error: ", err2)

	fmt.Println("Calling f2")
	val3, err3 := f2(32)
	fmt.Println(val3, " error: ", err3)
	val4, err4 := f2(42)
	fmt.Println(val4, " error: ", err4)

	_, e := f2(42)
	if ae, ok := e.(*argError); ok { // If you want to programmatically use the data in a custom error, you’ll need to get the error as an instance of the custom error type via type assertion.
		fmt.Println(ae.arg)
		fmt.Println(ae.problem)
	}
}
