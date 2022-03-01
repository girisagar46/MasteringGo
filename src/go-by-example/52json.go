package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"` // `json:"page"` means the json key is "page" even though field name is Page
	Fruits []string `json:"fruits"`
}

func main() {

	bolB, _ := json.Marshal(true) // Marshal returns the JSON encoding of v.
	fmt.Println(string(bolB))     // true

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB)) // 1

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB)) // 2.34

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB)) // ["apple","peach","pear"]

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// The JSON package can automatically encode your custom data types.
	// It will only include exported fields in the encoded output and will by default use those names as the JSON keys.
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B)) // {"Page":1,"Fruits":["apple","peach","pear"]}

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B)) // {"page":1,"fruits":["apple","peach","pear"]}

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil { // Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
		panic(err)
	}
	fmt.Println(dat) // map[num:6.13 strs:[a b]]

	num := dat["num"].(float64) // In order to use the values in the decoded map, we’ll need to convert them to their appropriate type. For example here we convert the value in num to the expected float64 type.
	fmt.Println(num)

	strs := dat["strs"].([]interface{}) // Accessing nested data requires a series of conversions.
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)           //{1 [apple peach]}
	fmt.Println(res.Fruits[0]) // apple

	// We can also stream JSON encodings directly to os.Writers like os.Stdout or even HTTP response bodies.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d) // {"apple":5,"lettuce":7}

}
