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
## var number int

To declare a variable called pi to be of type float64 with a value of 3.14159, you would write:
## var pi float64 = 3.14159

The value of a declared variable with no assignment will be its zero value.