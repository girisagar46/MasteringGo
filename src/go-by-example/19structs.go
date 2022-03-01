package main

import "fmt"

// Go’s structs are typed collections of fields. They’re useful for grouping data together to form records.

// Person struct type
type Person struct {
	name string // name field
	age  int    // age field
}

func buildPerson(name string, age int) *Person {
	p := Person{name: name, age: age}
	return &p // safely return a pointer to local variable as a local variable will survive the scope of the function.
}

func main() {
	fmt.Println(Person{"Scarlet", 26})           // create new struct
	fmt.Println(Person{name: "Bruce", age: 29})  // create new struct by specifying field name as well
	fmt.Println(Person{name: "Steve"})           // omitting a field will take a zero value depending on type
	fmt.Println(&Person{name: "Peter", age: 20}) // & prefix gets the pointer to the struct

	person1 := buildPerson("Steven", 50) //It’s idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(person1)
	fmt.Println(person1.name) // access a struct field with a dot operator

	tony := Person{name: "Tony Stark", age: 56}
	ironMan := &tony // ironMan is referencing tony

	ironMan.name = "Iron Man" // change in ironMan field will impact tony's field as well
	fmt.Println("ironMan", ironMan)
	fmt.Println("tony.name", tony.name)
}

/*
{Scarlet 26}
{Bruce 29}
{Steve 0}
&{Peter 20}
&{Steven 50}
Steven
ironMan &{Iron Man 56}
tony.name Iron Man
*/
