package main

import (
	"fmt"
	"unicode/utf8"
)

/*
NOTES:

A Go string is a read-only slice of bytes.
The language and the standard library treat strings specially - as containers of text encoded in UTF-8.
In other languages, strings are made of “characters”. In Go, the concept of a character is called a rune - it’s an integer that represents a Unicode code point.
More here: https://go.dev/blog/strings
*/
func main() {

	const ja = "こにちわ"             // string literals are UTF-8 encoded text.
	fmt.Println("Len: ", len(ja)) // Len:  12. Since strings are equivalent to []byte, this will produce the length of the raw bytes stored within.

	const en = "hello"
	fmt.Println("Len: ", len(en)) // Len:  5

	for i := 0; i < len(ja); i++ {
		fmt.Printf("%x ", ja[i])
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(ja)) // Output: 4
	// Note that the run-time of RuneCountInString depends on the size of the string, because it has to decode each UTF-8 rune sequentially.

	for idx, runeValue := range ja {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	var runeValue = 'こ' //Values enclosed in single quotes are rune literals.
	fmt.Println(runeValue)
	fmt.Printf("%#U\n", runeValue)

	// convert rune to string
	stringVal := string(runeValue)
	fmt.Println(stringVal)
}

/*
Len:  12
Len:  5
e3 81 93 e3 81 ab e3 81 a1 e3 82 8f
Rune count: 4
U+3053 'こ' starts at 0
U+306B 'に' starts at 3
U+3061 'ち' starts at 6
U+308F 'わ' starts at 9
12371
U+3053 'こ'
こ
*/
