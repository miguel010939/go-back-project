# GO Language

## Intro

Go is a statically-typed, compiled language with built-in concurrency support.
It's garbage-collected and memory safe.

Like Java, the garbage collector periodically deletes unreferenced memory. And unlike C, Go doesn't allow pointer arithmetic.

## Variables
### Primitive Types
Most of Go's primitive types are the same in other languages: integers (signed & unsigned), floating-point numbers, booleans & bytes (basically a char, or uint8).

In addition, strings are also a primitive type in Go, not a class of String objects or a vector of characters.
### Declaring Variables and Constants
Variables are declared with the keyword `var`. After the name of the variable comes the type. Multiple variables can be assigned or initialized in a single statement. 
When a variable is not initialized, it has a "zero value", which depends on its type. The common shorthand notation uses the symbol `:=` to declare a variable with the inferred type and assign it the value.


````go
var a string = "initial"
var b, c int = 1, 2
var d = true    // d is inferred to be boolean
var e int       // e has value 0
f := "apple"    // common shorthand
````
> [!TIP]
> The most common way to initialize a variable is using the shorthand, e.g. `x := 3.14`

Constants are defined analogously, using the keyword `const`.

### Pointers
The use of pointers is relatively limited, in comparison with other languages like C/C++. They are mainly used to work with references to values in functions, methods and data structures.

C-like syntax: `&` returns the memory address and `*` dereferences teh pointer, returning the value at the address.
```go
i := 1
ptr := &i

fmt.Println("ptr:", ptr)    // 0x42424242
fmt.Println("i:", *ptr)     // 1
```

> [!TIP]
> ```javascript
>//TODO **Automatic dereferencing**
>```
>  


## Builtin Data Structures
### Arrays
### Slices
### Maps

## Control Flow
Control flow in Go is quite similar to other programing languages. `Note on syntax`: parenthesis are not used around conditions
### Conditionals
Standard if-else block
```go
if x%2 == 0 {
	fmt.Println("x is even")
} else {
	fmt.Println("x is odd")
}
```
In Go you can write a statement to execute right before the condition is evaluated. It would make sense if it's somehow related
```go
if x := math.Cos(100.1); x > 0 {
	fmt.Println("x is positive")
}
```
Switch statements work very similarly to other languages
```go
switch i {
    case 1:
        fmt.Println("one")
    default:
        fmt.Println("something else")
}
```
### Loops
All loops use the keyword `for`. We can loop forever, while a condition is true, a given number of times
or over an array, slice, etc
```go
// equivalent to while(true)
for { 
	fmt.Println("This loops forever")
}
// while(condition)
for i <= 3 {
    fmt.Println(i)
    i = i + 1
}
// iteration over a slice
numbers := []int{1, 2, 3, 4, 5}
for i := range numbers {
    fmt.Println("range", i)
}
```

## Functions

## Errors
A salient feature of Go is that it treats errors as values instead of catching exceptions
### Errors as values
Functions that can *fail*, return an error type, instead of throwing an exception. If all goes well,
the error returned will be `nil` (null), otherwise it will have some value.

The compiler expects the programmer to handle error values and will protest if this is not done (at least syntactically). 
Errors are treated like any other variable. We can use the same common control flow structures (`if`, `switch`, `for`, etc).


### Error interface
Errors are type error, a built-in interface that implements `Error() string`.
This allows the programmer, among other things, to define custom error types that will, nonetheless, 
be recognized by the compiler as errors
```go
type MyError struct {
	Number  int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error number %d", e.Number)
}
```
## OOP? Kinda...
### Structs

```javascript
// TODO Public & Private fields in structs
```

### Methods
### Interfaces

## Concurrency
### Mutex
### Goroutines
### Channels


# Standard Library

## http/net
```javascript
// TODO Explain the main features I use from the package
```


## sql
```javascript
// TODO Explain the main features I use from the package
```

