package main

import "fmt"

func main() {
	nextInt := intSeq()
	fmt.Printf("Type of nextInt is: %T\n", nextInt) // Type of nextInt is: func() int
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	for {
		num := nextInt() // when this statement is executed, the return value is 3
		fmt.Println("num", num)
		if num > 5 {
			break
		}
	}

	fmt.Println("Reinitialize the intSeq")
	newNextInt := intSeq() // this function return has its own context with bounded i
	fmt.Println(newNextInt())
	fmt.Println(newNextInt())
}

func intSeq() func() int { // here intSeq returns anonymous function with return type int
	i := 0
	return func() int { // The returned function closes over the variable i to form a closure.
		i++
		return i
	}
}
