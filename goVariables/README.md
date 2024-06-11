# Go's basic variable types are:

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

# DECLARING A VARIABLE

Variables are declared using the var keyword. For example, to declare a variable called number of type int, you would write:
### var number int

To declare a variable called pi to be of type float64 with a value of 3.14159, you would write:
### var pi float64 = 3.14159

The value of a declared variable with no assignment will be its zero value.

# SHORT VARIABLE DECLARATION

nside a function (like the main function) the := short assignment statement can be used in place of a var declaration. The := operator infers the type of the new variable based on the value. It's colloquially called the walrus operator because it looks like a walrus... sort of.

These two lines of code are equivalent:

### var empty string
### empty := ""

### numCars := 10 // inferred as an integer
### temperature := 0.0 // temperature is inferred as a float because it has a decimal
### var isFunny = true // inferred as a boolean

Outside of a function (in the global/package scope), every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

