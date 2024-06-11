// SHORT VARIABLE DECLARATION

// Inside a function (like the main function) the := short assignment statement can be used in place of a var declaration.
// The := operator infers the type of the new variable based on the value.
// It's colloquially called the walrus operator because it looks like a walrus... sort of.

// var empty string this can be written as ------> empty := ""

// numCars := 10 // inferred as an integer
// temperature := 0.0 // temperature is inferred as a float because it has a decimal
// var isFunny = true // inferred as a boolean

package main

import "fmt"

func main() {
	// declare here
	messageStart := "Happy birthday! You are now"
	age := 21
	messageEnd := "years old!"

	// don't edit below this line
	fmt.Println(messageStart, age, messageEnd)
}
