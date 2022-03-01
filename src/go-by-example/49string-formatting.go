package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)
	// struct1: {1 2}

	fmt.Printf("struct2: %+v\n", p) // %+v variant will include the struct’s field names.
	// struct2: {x:1 y:2}

	fmt.Printf("struct3: %#v\n", p) // %#v variant prints a Go syntax representation of the value
	// struct3: main.point{x:1, y:2}

	fmt.Printf("type: %T\n", p) // %T prints the type of the value
	// type: main.point

	fmt.Printf("bool: %t\n", true) // Formatting booleans is straight-forward.
	// bool: true

	fmt.Printf("int: %d\n", 123) // %d for standard, base-10 formatting.
	// int: 123

	fmt.Printf("bin: %b\n", 14) // %b for binary formatting.
	// bin: 1110

	fmt.Printf("char: %c\n", 33) // %c for character formatting.
	// char: !

	fmt.Printf("hex: %x\n", 456) // %x for hexadecimal formatting.
	// hex: 1c8

	fmt.Printf("float1: %f\n", 78.9) // %f for floating-point formatting.
	// float1: 78.900000

	fmt.Printf("float2: %e\n", 123400000.0) // %e for scientific notation.
	// float2: 1.234000e+08

	fmt.Printf("float3: %E\n", 123400000.0) // %E for scientific notation.
	// float3: 1.234000E+08

	fmt.Printf("str1: %s\n", "\"string\"") // %s for string formatting.
	// str1: "string"

	fmt.Printf("str2: %q\n", "\"string\"") // %q for string formatting.
	// str2: "\"string\""

	fmt.Printf("str3: %x\n", "hex this") // %x for string formatting.
	// str3: 6865782074686973

	fmt.Printf("pointer: %p\n", &p) // %p for pointer formatting.
	// pointer: 0xc000014080

	fmt.Printf("width1: |%6d|%6d|\n", 12, 345) // %6d for right-justified decimal formatting.
	// width1: |    12|   345|

	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45) // %6.2f for left-justified decimal formatting.
	// width2: |  1.20|  3.45|

	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45) // %-6.2f for left-justified decimal formatting.
	// width3: |1.20  |3.45  |

	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b") // %6s for right-justified string formatting.
	// width4: |   foo|     b|

	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b") // %-6s for left-justified string formatting.
	// width5: |foo   |b     |

	s := fmt.Sprintf("sprintf: a %s", "string") // Sprintf formats and returns a string.
	fmt.Println(s)                              // sprintf: a string

	_, err := fmt.Fprintf(os.Stderr, "io: an %s\n", "error") //io: an error
	if err != nil {
		return
	}
}
