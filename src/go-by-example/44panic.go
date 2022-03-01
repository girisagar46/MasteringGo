package main

import "os"

/*
A common use of panic is to abort if a function returns an error value that we don’t know how to (or want to) handle.
Here’s an example of panicking if we get an unexpected error when creating a new file.
*/
func main() {
	// panic("a problem") // panic will exit the program hence code below panic is unreachable

	_, err := os.Create("/dasdsad/file")

	// Note that unlike some languages which use exceptions for handling of many errors, in Go it is idiomatic to use error-indicating return values wherever possible.
	if err != nil {
		panic("a problem")
	}
}

/*
panic: a problem

goroutine 1 [running]:
main.main()
        .../go-by-example/44panic.go:16 +0x74
*/
