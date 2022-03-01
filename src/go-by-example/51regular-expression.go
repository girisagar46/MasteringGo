package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach") // This tests whether a pattern matches a string.
	fmt.Println(match)                                     // true

	r, _ := regexp.Compile("p([a-z]+)ch") // Compile a regular expression pattern into a Regexp object.

	fmt.Println(r.MatchString("peach")) // Test whether a Regexp matches a string.

	fmt.Println(r.FindString("peach punch")) // Find the match for a Regexp within a string.

	fmt.Println("idx:", r.FindStringIndex("peach punch")) // Find the index sequence for a Regexp within a string.
	// idx: [0 5]

	fmt.Println(r.FindStringSubmatch("peach punch")) // Find the first match for a Regexp within a string.
	// [peach ea]

	fmt.Println(r.FindStringSubmatchIndex("peach punch")) // Find the index sequence for the first match for a Regexp within a string.
	// [0 5 1 3]

	fmt.Println(r.FindAllString("peach punch pinch", -1)) // Find all matches for a Regexp within a string. -1 means all matches.
	// [peach punch pinch]

	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1)) // Find all matches index for a Regexp within a string.
	// all: [[0 5 1 3] [6 11 7 9] [12 17 13 15]]

	fmt.Println(r.FindAllString("peach punch pinch", 2)) // Find up to n matches for a Regexp within a string.
	// [peach punch]

	fmt.Println(r.Match([]byte("peach"))) // Test whether a Regexp matches a byte slice.
	// true

	r = regexp.MustCompile("p([a-z]+)ch") // Compile a regular expression, panic on error.
	fmt.Println("regexp:", r)             // regexp: p([a-z]+)ch

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) // Replace the first match for a Regexp within a string with a replacement string.
	// a <fruit>

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper) //The Func variant allows you to transform matched text with a given function.
	fmt.Println(string(out))
}
