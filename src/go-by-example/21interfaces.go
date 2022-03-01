package main

import (
	"fmt"
	"math"
)

type geometry interface {
	calculateArea() float64
	calculatePerimeter() float64
}

type rectangle struct {
	width, length float64
}

type circle struct {
	radius float64
}

func (r rectangle) calculateArea() float64 {
	return r.length * r.width
}

func (r rectangle) calculatePerimeter() float64 {
	return 2 * (r.length + r.width)
}

func (c circle) calculateArea() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) calculatePerimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("Area:", g.calculateArea())
	fmt.Println("Perimeter:", g.calculatePerimeter())
}

func main() {
	r := rectangle{width: 3, length: 10}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
