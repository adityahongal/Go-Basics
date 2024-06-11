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

# TYPE INFERENCE

To declare a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.

When the right hand side of the declaration is typed, the new variable is of that same type:

#### var i int
### j := i          // j is also an int

However, when the right hand side is a literal value (an untyped numeric constant like 42 or 3.14), the new variable will be an int, float64, or complex128 depending on its precision:

#### i := 42           // int
#### f := 3.14         // float64
#### g := 0.867 + 0.5i // complex128

# SAME LINE DECLARATIONS

You can declare multiple variables on the same line:
#### mileage, company := 80276, "Tesla"

# TYPE SIZES

Ints, uints, floats, and complex numbers all have type sizes.

int  int8  int16  int32  int64 // whole numbers

uint uint8 uint16 uint32 uint64 uintptr // positive whole numbers

float32 float64 // decimal numbers

complex64 complex128 // imaginary numbers (rare)

The size (8, 16, 32, 64, 128, etc) represents how many bits in memory will be used to store the variable. The "default" int and uint types refer to their respective 32 or 64-bit sizes depending on the environment of the user.

The standard sizes that should be used unless the developer has a specific need are:

int
uint
float64
complex128

### Some types can be converted like this:
#### temperatureFloat := 88.26
#### temperatureInt := int64(temperatureFloat)

# WHEN SHOULD I USE A MORE SPECIFIC TYPE?

### When you're super concerned about performance and memory usage.

#### OR Unless you have a good reason to, stick to the following types:

- bool
- string
- int
- uint
- byte
- rune
- float64
- complex128

# CONSTANTS

Constants are declared with the const keyword. They can't use the := short declaration syntax.

#### const pi = 3.14159

# COMPUTED CONSTANTS

Constants must be known at compile time. They are usually declared with a static value:
#### const myInt = 15

For example, this is valid:

#### const firstName = "Lane"
#### const lastName = "Wagner"
#### const fullName = firstName + " " + lastName

That said, you cannot declare a constant that can only be computed at run-time like you can in JavaScript. 

#### This breaks:
// the current time can only be known when the program is running
#### const currentTime = time.Now()

# FORMATTING STRINGS IN GO

Go follows the printf tradition from the C language. In my opinion, string formatting/interpolation in Go is less elegant than Python's f-strings, unfortunately.

#### fmt.Printf - Prints a formatted string to standard out.
#### fmt.Sprintf() - Returns the formatted string

These following "formatting verbs" work with the formatting functions above:

### DEFAULT REPRESENTATION

#### The %v variant prints the Go syntax representation of a value, it's a nice default.
s := fmt.Sprintf("I am %v years old", 10)
// I am 10 years old

s := fmt.Sprintf("I am %v years old", "way too many")
// I am way too many years old

### STRING

s := fmt.Sprintf("I am %s years old", "way too many")
// I am way too many years old

### INTEGER

s := fmt.Sprintf("I am %d years old", 10)
// I am 10 years old

### FLOAT

s := fmt.Sprintf("I am %f years old", 10.523)
// I am 10.523000 years old

// The ".2" rounds the number to 2 decimal places

s := fmt.Sprintf("I am %.2f years old", 10.523)
// I am 10.52 years old

