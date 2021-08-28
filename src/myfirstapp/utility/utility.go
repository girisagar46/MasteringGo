package utility

import (
    "fmt"
)

func SayHi() {
    fmt.Println("Hi!")
}

// sayHello is not exportable function because the first letter of the function name is in lowercase
// this function is private function
func sayHello() {
    fmt.Println("Hello!")
}

// we do not need to write (a int, b int) as go understands both are int if we
// just specify once
// Similarly, here return type is (int, int, int) indicating multiple values will
// be returned by this function and all return type is of int
/*
func mathematics(a, b int) (int, int, int) {
    return a + b, a-b, a*b
}
*/

// This declaration is more verbose and clear.
// Here, in return type, we define variable names, and we just use the return keyword
func mathematics(a, b int) (addition, subtraction, multiplication int) {
    addition = a + b
    subtraction = a - b
    multiplication = a * b
    return
}

// ComputeAddAndMultiply function takes function as an argument
// and execute it inside the function
// here, `fn` is a function which matches the signature of `add` function
func ComputeAddAndMultiply(val int, fn func(a, b int) int) int {
    return val * fn(val, val)
}

func Inner() (res, computeRes int){
    // This is a function assigned to a variable (Function values)
    // These type of function can be only defined inside another function
    // See more: https://tour.golang.org/moretypes/24
    add := func(a, b int) int {
        return a+ b
    }
    // calling internally
    res = add(1, 2)
    computeRes = ComputeAddAndMultiply(5, add)
    return
}

// Increment closures
// Here inner function is bound to tha variable and we're returning function
func Increment() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

// ChangeMe do not need to return anything, because we're directly working with the address
func ChangeMe(X *int) {
    *X = 10
}
