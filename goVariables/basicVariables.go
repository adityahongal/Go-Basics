// # DECLARING A VARIABLE

// Variables are declared using the var keyword.
// For example, to declare a variable called number of type int, you would write:
// var number int

// To declare a variable called pi to be of type float64 with a value of 3.14159, you would write:
// var pi float64 = 3.14159

package main

import "fmt"

func main() {
	// initialize variables here
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
