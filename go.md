# `Go short summary`

## `Variables`

### Declaration

```go
var foo int
var foo int = 42
// Short code
foo := 42
```

- Can't redeclare variables, but can shadow them
- All variables must be used. Else it will throw compilation error

### Visibility

- Lower case first letter for package scope
- Upper case first letter to export
- no private scope

### Naming convention

- Pascal case or camel case
- Capitalize acronyms
- As short as reasonable
- Longer names for longer lives

### Type conversions

- Destination Types(variable)
- Use strconv package for strings

## `Primitive Data Types`

    - Boolean Type
    - Numeric Type
        - Integers
        - Floating point
        - Complex Numbers
    - Text Type
        - string
        - rune

### Boolean Type

#### Declaration

```go
var isHidden bool = true
isHidden := false
```

- Values are `true` or `false`
- Not an alias for other types.
- zero value is `false`

### Nummeric Types

#### Declaration

```go
var i int = 32
j := 22
var k float32 = 23.4
var l complex64 = 23 + 12i

```

- Integers
  - Signed Integers
    - inte type has varying but min 32 bits
    - 8 bit(int8) through 64 bit(int64)
  - Unsigned Integers
    - 8 bit(byte and uint8) through 32 bit(uint32)
  - Arithmatic operations
    - Addition, subtraction, multiplication, division, remainder
  - Bitwise operations
    - And, or, xor and not
  - Zero value is 0
  - Can't mix types in same family!(uint16 + uint32 = error)
- Floating point Numbers
  - Follow IEEE-754 standard
  - Zero value is 0
  - 32 and 64 bit versions
  - Literal styles
    - Decimal(3.14)
    - Exponential(13e18 or 2E10)
    - Mixed(13.4e23)
  - Arithmatic operations
    - Addition, subtraction, multiplication, division
- Complex Numbers
  - Zero value is (0 + 0i)
  - 64 and 128 but versions
  - Built-in functions
    - complex - make complex number from two floats
    - real - get real part as float
    - imag - get imaginary part as float
  - Arithmatic operations
    - Addition, subtraction, multiplication, division

### Text types

#### Declaration

```go
// String type
var s string = "this is a string"
s := "This is a string"
// Type rune
var str rune = "This is a rune"
str := 'this is also a rune'
```

- Strings
  - UTF-8
  - Immutable
  - Can be concatenated with plus (+) operator
  - Can be converted to []byte
- Rune
  - UTF-32
  - Alias for int32
  - SPecial methods normally required to process

## `Constants`

### Declaration

```go
const myConst string = "some constant"
```

- Immutable, but cannot be shadowed
- Replaced by the compiler at compile time
  - Value must be calculable at compile time
- Named like variables
  - PascalCase for exported constants
    = CamelCase for internal constants
- Typed constants work like immutable variables
  - Can interoperate only with the same type
- Untyped constants work like literals
  - Can interoperate with similar types.
- Enumarated constants
  - Special symbol `iota` allows related constants to be created easily
  - `iota` starts at 0 in each const block and increments by one.
  - Watch out of constant values that match zero values for variables
- Enumerated expressions
  - Operations that can be determined at compile time are allowed
    - Arithmetic
    - Bitwise operations
    - Bitshifting

### Example

```go
package main

import (
	"fmt"
)

const (
	isAdmin = 1 << iota
	isHeadQuarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {

	var roles = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Println(roles)
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is admin? %v\n", isAdmin & roles == isAdmin)
	fmt.Printf("Is Head Quarters? %v\n", isHeadQuarters & roles == isHeadQuarters)
}
```

## `Arrays and slices`

### Example
```
package main

import (
	"fmt"
)

func main() {
	// grades := [3]int{97, 85, 93 }
	var students [3]string
	students[0] = "Lisa"
	students[1] = "Arnold"
	grades := [...]int{97, 85, 93}
	fmt.Printf("Grades: %v, %T", grades, grades)
	fmt.Printf("Students: %v, %T", students, students)
	fmt.Printf("No of Students: %v, %T\n", len(students), students)

	// Array of arrays
	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)

	// Slice
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]   // Slice of all elements
	c := a[3:]  // Slice from 4th element to the end
	d := a[:6]  // Slice first 6 elements
	e := a[3:6] // Slice the 4th, 5th and 6th elements

	// Slices are natuarally passed by reference
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	// Built-in make function
	// type len cap
	// x := make([]int, 3, 100)
	x := []int{1,2}
	x = append(x, 1)
	x = append(x, []int{1, 1, 1, 1, 1, 1}...)
	fmt.Println(x)
	// Remove first element
	y := a[1:]
	// Remove the last element
	z := a[:len(a) -1]
	// Remove an element from a particular index
	w := append(a[:2], a[5:]...)
	fmt.Printf("Length: %v\n", y)
	fmt.Printf("index sliced: %v\n", w)
	fmt.Printf("Capacity: %v\n", z)

}
```

#### Arrays
- collection of items with same type
- Fixed size
- Access via zero-based index
- len function returns size of array
- Copies refer to different underlying data

#### Slices
- Backed by array
- make function can be used to pre-define the capacity of the slice
- len & cap functions determine the lenght and capacity of the slice respectively.
- append function to add elements to the slice
    - May cause expensive copy operation if underlying array is too small
- Copies refer to same underlying array

## `Maps and Structs`

### Example
```go
package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

// embedding 
type Animal struct {
	Name string `required max: "100"`
	Origin string
}

type Bird struct {
	Animal	//Embedding animal type
	SpeedKPH float32
	canFly bool
}


func main() {
	statePopulations := map[string]int{
		"california": 39250017,
		"Texas":      27862596,
	}
	// add to  map
	statePopulations["Georgia"] = 89789900
	fmt.Println(statePopulations)

	// Delete an item from map
	delete(statePopulations, "california")

	_, ok := statePopulations["Oho"]
	fmt.Println(ok)

	countryPopulation := make(map[string]int)

	fmt.Println(countryPopulation)

	fmt.Println(len(statePopulations))

	// Struct

	aDoctor :=
		Doctor{
			number:    3,
			actorName: "John Cena",
			companions: []string{
				"Liz Shaw",
				"Jo Grant",
			},
		}

	fmt.Println(aDoctor)
	
	anotherDoctor := aDoctor
	anotherDoctor.actorName = "Tom Baker"
	
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)
	
	// anonymous struct 
	someStruct := struct{name string}{name: "John Pertwee"}
Learn Go Programming - Golang Tutorial for Beginners

	fmt.Println(someStruct)
	
	// Composition through embedding
	b := Bird {
		Animal: Animal{ Name: "Emu", Origin: "Australia"},
	}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.canFly = false
	
	fmt.Println(b)
	
	// Tags
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}

```

### Maps
- collections of value types that are accessed via keyys
- Created via literals or via make functon
- Members accessed via [key] syntax
    - myMap["key"] = "value"
- Check for presence with "value, ok" form of result
- Multiple assignments refer to same underlying data.(reference type)

### Structs
- COllections of disparate data tyypes that describe a single concept
- Keyed by name fields
- Normally created as types, but anonymous structs are allowed
- Structs are value types
- No inheritance, but can use composition via embedding
- Tags can be added to struct fields to describe field

## `Control Flow`

### If statements

#### Example
```go
package main

import (
	"fmt"
	"math"
)

func main() {

	statePopulations := map[string]int{
		"california": 39250017,
		"Texas":      27862596,
	}

	// No single line block
	if true {
		fmt.Println("The test is true")
	}
	// can add an initialiser
	if pop, ok := statePopulations["Texas"]; ok {
		fmt.Println(pop)
	}

	number := 50
	guess := 30
	if guess < 1 || returnTrue() || guess > 100 {
		fmt.Println("The guess must be between 1 and 100")
	} else {
		if guess < number {
			fmt.Println("Too low")
		}

		if guess > number {
			fmt.Println("Too high")
		}

		if guess == number {
			fmt.Println("You got it")
		}
		fmt.Println(number <= guess, number >= guess, number != guess)
	}
	
	myNum := 0.123
	if math.Abs(myNum /math.Pow(math.Sqrt(myNum), 2) - 1) < 0.001 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}
	
	switch i := 2 + 3;i {
		case 1, 3, 10:
			fmt.Println("One three ten")
		case 2, 4, 9:
			fmt.Println("Two four nine")
		default:
			fmt.Println("default")
	}
	j := 10
	switch {
		case j >= 20:
			fmt.Println("passed greater than 20")
			// fallThrough is logic less. will fallthrough to the next case irrespective of the case
		case j <= 20:
			fmt.Println("passed less than 20") 
	}
	
	var i interface{} = 1
	switch i.(type) {
		case int:
			fmt.Println("i is an integer")
			// break can be used
		case string:
			fmt.Println("i is a string")
		default:
			fmt.Println("i is another type")
	}

}

func returnTrue() bool {
	fmt.Println("Return true")
	return true
}

```
#### If statement
- Initializer
- Comparison operators
- Logical operators
- Short circuiting
- If - else statement
- If - else if statement
- Equality and floats
    - Need an error param

#### Switch statement
- Switching on a tag
- Cases with multiple tests
- Initializers
- Switch with no tag
- Fallthrough
    - Implicit break
    - Explicit fallThrough
- Type Switches
- Breaking out early


## `LOOPING`

### For statement

- Simple loops
	- for initializer; test; incrementer {}
	- for test {}
	- for {}
- Exiting early
	- break
	- continue
	- labels
- Loop over collections
	- arrays, slices, maps, channels
	- for k, v := range collection {}


### Defer
#### Example

```go
package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Start")
	// defer executes just before the complete function exits just before the return
	// LIFO order
	defer fmt.Println("middle")
	defer fmt.Println("end")
	// Panic happens after defered statements are executed
	panic("something happened")
	

	a := "start"
	// Function arguments get assigned as usual
	defer fmt.Println(a)
	a = "end"
	
	// panic
	// x, y := 1, 0
	// ans := x / y
	fmt.Println("start")
}
```
- Used to delay the execution of a statement until function exits
- Useful to group "open" and "close" functions together
	- Be careful in loops
- Run in LIFO order
- Argumentss evaluated at the time defer is executed, not at the time of called function execution

### Panic
#### Example
```go
package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			panic(err)
		}
	}()
	panic("Something bad happened")
	fmt.Println("done panicking")
}
```
- Occur when program cannot continue at all
	- Don't use when file can't be opened, unless it is critical
	- Use for unrecoverable events - cannot obtain TCP port for webserver
- Function will stop executing
	- Deferred functions will still fire
- If nothing handles panic, program will exit

### Recover
- Used to recover from panics
- Only useful in deferred functions
- Current function will not attempt to continue, but higher functions in call stack will continue to propogate

## `Pointers`

	A slice contains a pointer to the first element of an underlying array. So when we copy or assign a slice to another variable, it is actually copying the pointers to the underlying array. It's the same at the case of a map.

### Example
```go
package main

import (
	"fmt"
)

type myStruct struct {
	foo int
}
func main() {
	var a int = 42
	// Point to a
	var b *int = &a
	fmt.Println(a, b)
	// Print address of a 
	fmt.Println(&a)
	a = 27
	fmt.Println(a, *b)
	
	// pointer arithmetics
	x := [3]int{1,2,3}
	y := &x[0]
	z := &x[1]
	fmt.Printf("%v %p %p\n", x, y, z)
	// No pointer arithemetics
	
	
	// structs pointers
	var ms *myStruct
	ms = new(myStruct)
	// ms = &myStruct{foo: 42}
	
	ms.foo = 42
	fmt.Println(ms.foo)
	
}
```

- Creating pointers
	- Pointer types use an asterisk(*) as a prefix to type pointed to
		- *int - a pointer to an integer
	- Use the addressof operator (&) to get the address of variable
- Dereferencing pointers
	- Dereference a pointer by preceding with an asterisk(*)
	- Complex types(e.g. structs) are automatically dereferenced
- Create pointers to objects
	- Can use the address of operator (&) if value type already exists
		- ms := myStruct{foo: 42}
		- p := &ms
	- Use addressof operator before initializer
		- &myStruct{foo: 42}
	- Use the new keyword
		- Can't initialize fields at the same time
- Types with internal pointers
	- All assignment operations in Go are copy operations
	- Slices and maps contain internal pointers, so copies point to same underlying data.


## `Functions`
### Syntax
```go
func main() {
	// Somthting
}
```

### Example
```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	greeting := "Hello"
	name := "Stacey"
	sayMessage(greeting, &name)
	fmt.Println(name)
	
	// anonymous function
	func() {
		fmt.Println("Hello Go")
	}()
	
	var f func() = func() {
		fmt.Println("Hello Go!")
	}
	f()
	
	var dividefunc func(float64, float64) (float64, error)
	dividefunc = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Cannot divide by zero")
		} else {
			return a /b, nil
		}
	}
	quotient, err := dividefunc(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(quotient)
	
	// Anonymous functions can be really useful in the case of async operations inside a loop
	
	// Methods
	
	g := greeter {
		greeting: "Hello",
		name: "Go",
	}
	g.greet()
}

type greeter struct {
	greeting string
	name string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
func sayMessage(msg string, name *string) {
	fmt.Println(msg, *name)
	*name = "Ted"
	fmt.Println(name)
	s := sum("The sum is", 1,2,3,4,5,6)
	fmt.Println(s)
	d, err :=divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

//func sum(msg string, values ...int) *string {
//	fmt.Println(values)
//	result := 0
//	for _, v := range values {
//		result += v
//	}
//	concatenated := strings.Join([]string{msg, strconv.Itoa(result)}, " ")
//	return &concatenated
//}
func sum(msg string, values ...int) (concatenated string) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	concatenated = strings.Join([]string{msg, strconv.Itoa(result)}, " ")
	return
}

func divide(a,b float64) (float64, error) {
	if b== 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
```

### parameters
- Comma delimeted list of variables and types
	- func foo(bar string, baz int)
- Parameters of same type list type once
	- func foo(bar, baz int)
- When pointers are passed in, the function can change the value in the caller
	- This is always true for data of slices and maps\
- User variadic parameters to send list of same types in
	- Must be last parameter
	- Recieved as a slice
	- func foo(bar string, baz ...int)

### Return values
- Single return values just list type
	- func foo() int
- Multiple return value list types surrounded by parantheses
	- func foo() (int, error)
	- The (result type, error) paradigm is a very common idiom
- Can use named return values
	- Initializes returned variables
	- Return using return keyword on its own
- Can return the address of local variables
	- Automatically promoted from local memory(stack) to shared memory(heap)

### Anonymous functions
- Functions don't have names if they are:
	- Immediately invoked
	```go
	func() {
		...
	}()
	```
	- Assigned to a variable or passed as an argument to a function
	```go
	a := func() {
		...
	}
	a()
	```

### Functions as types
- Can assign functions to variables or use as arguments and return values in functions
- Type signature is like function signature, with no parameter names
	```go
	var f func(string, string, int) (int, error)
	```

### Methods
- Function that executes in context of a type
- Format
	```go
	func(g greeter) greet() {
		...
	}
	```
- Reciever can be value or pointer
	- Value reciever gets copy of type
	- Pointer reciever gets pointer to type



## `Interfaces`
An interface is a set of method signatures. When a type provides definition for all the methods in the interface, it is said to implement the interface.

	For example WashingMachine can be an interface with method signatures Cleaning() and Drying(). Any type which provides definition for Cleaning() and Drying() methods is said to implement the WashingMachine interface.

<!-- It's probably something similar to inheritance -->
	It is legal to call a pointer-valued method on anything that is already a pointer or whose address can be taken. The concrete value stored in an interface is not addressable and hence it is not possible for the compiler to automatically take the address of a in line no. 45 and hence this code fails.


## `Go Routines`
- Cheap compared to threads. Only a few kb in stack size
- Can go grow or shrink according to the need of the program. But thread size in a stack has to be fixed.

- The Goroutines are multiplexed to fewer number of OS threads. There might be only one thread in a program with thousands of Goroutines. If any Goroutine in that thread blocks say waiting for user input, then another OS thread is created and the remaining Goroutines are moved to the new OS thread. All these are taken care by the runtime and we as programmers are abstracted from these intricate details and are given a clean API to work with concurrency.
- Goroutines communicate using channels. Channels by design prevent race conditions from happening when accessing shared memory using Goroutines. Channels can be thought of as a pipe using which Goroutines communicate. We will discuss channels in detail in the next tutorial.



	
