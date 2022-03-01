package main

import (
	"fmt"
	m "math" // import math aliased as m
)

// const keyword is used to declare constant
// const can appear anywhere a var can appear
const s string = "constant"

var t string = "time" // we can omit the type here as it's implicit

func main() {
	fmt.Println(s)
	fmt.Println(t)

	// s = "123" // this results an error as we cannot re-assign const
	t = "clock" // variable can be mutated with the same type
	// t = 34      // variable can not be mutated with other type this will result error. "TYPE IS LIFE" in go
	fmt.Println(t)

	const p = 5555555
	fmt.Println(p)

	const n = 5_555_555 // we can use underscore to make large numbers readable
	fmt.Println(n)

	const d = 3e20 / n // Constant expressions perform arithmetic with arbitrary precision.
	fmt.Println(d)

	fmt.Println(m.Sin(n))

}

/**
constant
time
clock
5555555
5555555
5.400000540000054e+13
0.24789192322632397
*/
