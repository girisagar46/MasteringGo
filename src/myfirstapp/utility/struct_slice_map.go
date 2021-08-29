package utility

import (
	"fmt"
)

type CMember struct {
	name      string
	age       int
	address   string
	rank      string
	clearance int
}

// PrintSecurityClearance method
func (cm1 CMember) PrintSecurityClearance() {
	fmt.Println(cm1.clearance)
}

func StructSliceMap() {
	fmt.Println()
	fmt.Println("====STRUCT AND SLICE START===")
	structAndSliceDemo()
	fmt.Println("====STRUCT AND SLICE END===")

	fmt.Println("====MAPS START===")
	mapsDemo()
	fmt.Println("====MAPS END===")
}

func mapsDemo() {
	// NOTE: Maps are not thread safe
	//var m map[string]CMember
	m := make(map[string]CMember)
	// If key does not exist, it will be created and value will be assigned.
	m["SG"] = CMember{"Sagar1", 26, "Station Mars", "Lieutenant", 5}
	fmt.Println(m)

	avengers := map[string]CMember{
		"tony":            {"Tony Stark", 45, "Station Earth", "Admin", 10},
		"captain_america": {"Steve Rogers", 49, "Station Earth", "Admin", 9},

	}
	avengers["black_widow"] = CMember{"Natasha Romanoff", 30, "Station Earth", "Admin", 5}
	fmt.Println(avengers)

	// accessing elements
	natasha := avengers["black_widow"]
	fmt.Println(natasha)

	v, ok := avengers["black_widow"] // check if key, value pair exists in a map
	fmt.Println(v, ok) // if key, value pair does not exist, it'll print `{ 0   0} false`

	// deleting an key, value entry
	delete(avengers, "black_widow")
	fmt.Println(len(avengers))

	// looping over map
	for k, v := range avengers {
		fmt.Println("key:", k, " value:", v)
	}
}

func structAndSliceDemo() {
	// declaring a struct
	cm1 := CMember{"Sagar1", 26, "Station Mars", "Lieutenant", 5}
	fmt.Println("cm2 direct: ", cm1)

	// calling the method
	cm1.PrintSecurityClearance() // prints 5

	cm2 := CMember{
		name:      "Sagar2",
		age:       26,
		address:   "Station Mars",
		rank:      "Lieutenant",
		clearance: 5,
	}
	fmt.Println("cm2 using key: ", cm2)

	var cm3 CMember
	cm3.name = "Sagar3"
	cm3.age = 26
	cm3.address = "Station Mars"
	cm3.rank = "Lieutenant"
	cm3.clearance = 5
	fmt.Println("cm3 using key: ", cm3)

	cmp := &cm3
	cmp.clearance = 4 // this is going to change cm3
	fmt.Println(cmp)

	var crew []CMember // Array of struct type
	crew = append(crew, cm1, cm2, cm3, CMember{"Joe", 32, "Station Earth", "Soldier", 1})
	fmt.Println(crew)

	// Slicing Arrays
	// Sub slices are reference to original array. Any change in element in sub slice will affect original array
	subCrew1 := crew[1:3] // Index 1 and 2. Here 3 is exclusive
	fmt.Println(subCrew1)

	subCrew2 := crew[:3] // Index 0, 1, 2.
	fmt.Println(subCrew2)

	subCrew3 := crew[1:] // All elements after the first one.
	fmt.Println(subCrew3)

	subCrew3[0].age = 100 // This will change the original array element.
	fmt.Println(crew)

	a := make([]int, 10) // make slice with slice type ([]int) and slice size (10)
	fmt.Println("a", a)

	for i, v := range crew {
		fmt.Println(i, v)
	}
}
