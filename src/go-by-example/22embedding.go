package main

import "fmt"

// Go support embedding of structs and interfaces to express a more seamless composition of types.

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type describer interface {
	describe() string
}

type container struct {
	base  // A container embeds a base. An embedding looks like a field without a name.
	value string
}

func main() {
	co := container{ // When creating structs with literals, we have to initialize the embedding explicitly
		base: base{
			num: 100,
		},
		value: "embedding",
	}

	fmt.Printf("cp={num: %v, name: %v}\n", co.num, co.value) // we can access the base field directly from co
	fmt.Println("also num: co.base.num ", co.base.num)       // we can also use full path to get the field value
	fmt.Println("describe: co.describe ", co.describe())     // method from base is also present in the container
	fmt.Println("describe: co.base.describe ", co.base.describe())

	var d describer = co // now the co also has describe method, container can implement the describer interface
	fmt.Println("describer:", d.describe())
}

/*
cp={num: 100, name: embedding}
also num: co.base.num  100
describe: co.describe  base with num=100
describe: co.base.describe  base with num=100
describer: base with num=100
*/
