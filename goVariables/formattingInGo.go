// FOrmatting in Go

// Create a userLog variable on line 15. It should contain:

// "Name: FNAME LNAME, Age: AGE, Rate: MESSAGERATE, isSubscribed: ISSUBSCRIBED, Message: MESSAGE"

// Where FNAME LNAME AGE MESSAGERATE ISSUBSCRIBED and MESSAGE correspond to the variables above.

// MESSAGERATE should be rounded to the tenths place.

// fmt.Sprintf can be used to format strings.
// %.1f rounds a float to the tenths place, %.2f rounds to the hundredths place, etc.
// %t formats a boolean value.
// %v can be used to format any value in its default representation.
// %s can be used to format a string.
// %d can be used to format an integer.

package main

import "fmt"

func main() {
	fname := "Darshan"
	lname := "Dboss"
	age := 45
	messageRate := 0.5
	isSubscribed := false
	message := "Sometimes a hypocrite is nothing more than a man in the process of changing."

	// Don't touch above this line

	userLog := fmt.Sprintf("Name: %v %v, Age: %d, Rate: %.1f, isSubscribed: %t, Message: %s",
		fname,
		lname,
		age,
		messageRate,
		isSubscribed,
		message)

	// Don't touch below this line

	fmt.Println(userLog)
}
