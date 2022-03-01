package main

import "fmt"

type MyInterface interface {
	methodWithoutParam()
	methodWithParam(float64)
	methodWithReturn() string
}

type MyType string

func (t MyType) methodWithoutParam() {
	fmt.Println("Method without parameter is called")
}
func (t MyType) methodWithParam(val float64) {
	fmt.Printf("Method with parameter is called with param: %v\n", val)
}
func (t MyType) methodWithReturn() string {
	fmt.Printf("Method with retrun type is called")
	return "Hello World"
}
func (t MyType) methodNotInInterface() string {
	fmt.Printf("This method is not in interface.")
	return "Hello World"
}

func main() {
	var value MyInterface
	value = MyType("Hello")
	value.methodWithoutParam()
	value.methodWithParam(45.70)
	value.methodWithReturn()
}
