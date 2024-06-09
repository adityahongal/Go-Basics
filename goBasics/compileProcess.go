//THE COMPILATION PROCESS

// Computers need machine code, they don't understand English or even Go.
// We need to convert our high-level (Go) code into machine language,
// which is really just a set of instructions that some specific hardware can understand.

// package main lets the Go compiler know that we want this code to compile and run as a standalone program.
// import fmt imports the fmt (formatting) package. The formatting package exists in Go's standard library
// and lets us do things like print text to the console.
// func main() defines the main function.
// main is the name of the function that acts as the entry point for a Go program.

//Run the program using --> go run filename.go

package main

import "fmt"

func compileProcess() {
	fmt.Println("the compiled textio server is starting")
}
