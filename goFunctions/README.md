# FUNCTIONS

Functions in Go can take zero or more arguments.

### To make code easier to read, the variable type comes after the variable name.

For example, the following function:

#### func sub(x int, y int) int {
####  return x-y
#### } 

Accepts two integer parameters and returns another integer.

#### Here, func sub(x int, y int) int is known as the "function signature".

# MULTIPLE PARAMETERS

When multiple arguments are of the same type, and are next to each other in the function signature, the type only needs to be declared after the last argument.

Here are some examples:

#### func addToDatabase(hp, damage int, name string) {
####  // ?
#### }

#### func addToDatabase(hp, damage int, name string, level int) {
####  // ?
#### }

# GO-STYLE SYNTAX

Go's declarations are clear, you just read them left to right, just like you would in English.

#### x int
#### p *int
#### a [3]int

# IGNORING RETURN VALUES

A function can return a value that the caller doesn't care about. We can explicitly ignore variables by using an 
#### underscore: _

For example:

func getPoint() (x int, y int) {
  return 3, 4
}

// ignore y value
#### x, _ := getPoint()

Even though getPoint() returns two values, we can capture the first one and ignore the second.

### WHY MIGHT YOU IGNORE A RETURN VALUE?

Maybe a function called getCircle returns the center point and the radius, but you only need the radius for your calculation. In that case, you can ignore the center point variable.

#### The Go compiler will throw an error if you have any unused variable declarations in your code, so you need to ignore anything you don't intend to use.