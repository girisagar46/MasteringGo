package main

import "fmt"

// Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages).
func main() {
	m := make(map[string]int) //make(map[key-type]val-type)
	fmt.Println(m)            // map[]

	// setting key-value pairs
	m["key1"] = 100
	m["key2"] = 200
	fmt.Println(m)

	// getting value via key
	v1 := m["key2"]
	fmt.Println(v1)
	fmt.Println("map length: ", len(m)) // len can be used in the map as well

	// deleting key-value pair
	delete(m, "key2")      // delete based on key. delete is build-in function
	delete(m, "key100000") // this has no effect as specified key is not in the map
	fmt.Println("After value deletion: ", m)

	val, present := m["key3"] // second return is optional which is a bool to indicate if specified value is present or not
	fmt.Println("is value present?: ", present)
	fmt.Println("but, the value is: ", val) // the value returned is 0. This might be confusing is map has 0 as value

	_, available := m["key100"] // we can use _ to ignore if we just want to check if key-value pairs are present
	fmt.Println("is value present?: ", available)

	prePreparedMap := map[string]int{"hello": 1, "world": 2} // declare and initialize a new map
	fmt.Println(prePreparedMap)
}

/*
map[]
map[key1:100 key2:200]
200
map length:  2
After value deletion:  map[key1:100]
is value present?:  false
but, the value is:  0
is value present?:  false
map[hello:1 world:2]
*/
