package main

import "fmt"

// range iterates over elements in a variety of data structures.
func main() {
	nums := []int{1, 2, 3, 4} // define and initialize a slice
	sum := 0
	for _, num := range nums { // loop over slice to sum up each value
		sum += num
	}
	fmt.Println(sum)

	for i, num := range nums { // range over slice/array provides index(i) and value(num)
		if num == 3 {
			fmt.Println("index = ", i)
		}
	}

	myMap := map[string]string{"a": "apple", "b": "ball"} // define and initialize a map[string]string
	for k, v := range myMap {                             // range over maps provides key, value pair
		fmt.Printf("key:%s -> value:%s\n", k, v)
	}

	for k := range myMap { // if we only want key, we can ignore the second property
		fmt.Println("key", k)
	}

	// range on strings iterates over Unicode code points.
	//The first value is the starting byte index of the rune and the second the rune itself.
	// Note that the starting byte index between unicode and string are different
	fmt.Println("Range over `こにしわ` (UNICODE)")
	for i, c := range "こにしわ" {
		fmt.Printf("index:%v, character:%c\n", i, c)
	}

	fmt.Println("Range over `hello` (ASCII)")
	for i, c := range "hello" {
		fmt.Printf("index:%v, character:%c\n", i, c)
	}
}

/*
10
index =  2
key:a -> value:apple
key:b -> value:ball
key a
key b
Range over `こにしわ` (UNICODE)
index:0, character:こ
index:3, character:に
index:6, character:し
index:9, character:わ
Range over `hello` (ASCII)
index:0, character:h
index:1, character:e
index:2, character:l
index:3, character:l
index:4, character:o
*/
