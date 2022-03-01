package main

import "fmt"

type rect struct {
	length, width int
}

// func (receiver type) funcName returnType {}
// this is a method with pointer receiver type
// pointer receiver is preferred inorder to avoid copying on method calls
func (r *rect) area() int { // This area method has a `receiver type` of *rect.
	return r.width * r.length
}

// this is a method with value receiver type
func (r rect) perimeter() int {
	return 2 * (r.width + r.length)
}

func main() {
	r := rect{10, 20}
	fmt.Println(r.area()) // we use dot operator to call struct methods
	fmt.Println(r.perimeter())

	rp := &r
	fmt.Println(rp.area())
	fmt.Println(rp.perimeter())
}

/*
200
60
200
60
*/
