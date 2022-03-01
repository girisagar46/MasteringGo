package main

import (
	"fmt"
	str "strings"
)

var p = fmt.Println // alias for fmt.Println; print with newline

func main() {

	p("Contains:  ", str.Contains("test", "es"))        // true
	p("Count:     ", str.Count("test", "t"))            // 2
	p("HasPrefix: ", str.HasPrefix("test", "te"))       // true
	p("HasSuffix: ", str.HasSuffix("test", "st"))       // true
	p("Index:     ", str.Index("test", "e"))            // 1
	p("Join:      ", str.Join([]string{"a", "b"}, "-")) // a-b
	p("Repeat:    ", str.Repeat("a", 5))                // aaaaa
	p("Replace:   ", str.Replace("foo", "o", "0", -1))  // f00
	p("Replace:   ", str.Replace("foo", "o", "0", 1))   // f0o
	p("Split:     ", str.Split("a-b-c-d-e", "-"))       // []string{"a", "b", "c", "d", "e"}
	p("ToLower:   ", str.ToLower("TEST"))               // test
	p("ToUpper:   ", str.ToUpper("test"))               // TEST
}
