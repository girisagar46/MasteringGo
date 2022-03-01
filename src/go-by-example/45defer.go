package main

import (
	"fmt"
	"os"
)

/*
Defer is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup.
defer is often used where e.g. ensure and finally would be used in other languages.
*/
func main() {
	f := createFile("./defer.txt")
	defer closeFile(f) // No matter where we write defer in code, this will be executed at the end of the enclosing function (main), after writeFile has finished.
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("Writing")
	_, err := fmt.Fprintln(f, "data")
	if err != nil {
		return
	}
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil { // It’s important to check for errors when closing a file, even in a deferred function.
		_, err := fmt.Fprintf(os.Stderr, "error: %v\n", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}

/*
creating
Writing
closing
*/
