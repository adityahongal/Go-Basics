// GO IS STRONGLY TYPED

// Go enforces strong and static typing, meaning variables can only have a single type.
// A string variable like "hello world" can not be changed to an int, such as the number 3.

// CONCATENATING STRINGS

// Two strings can be concatenated with the + operator.
// Because Go is strongly typed, it won't allow you to concatenate a string variable with a numeric variable.

package main

import "fmt"

func concatStrings() {
	var username string = "presidentSkroob"
	var password string = "12345"

	fmt.Println("Authorization: Basic", username+":"+password)
}

func main() {
	concatStrings()
}
