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
>  //TODO **Automatic dereferencing**


## Builtin Data Structures
### Arrays
### Slices
### Maps

## Control Flow
### Conditionals
### Loops

## Functions

## Errors
### Errors as values
### Error interface

## OOP? Kinda...
### Structs
### Methods
### Interfaces

## Concurrency
### Mutex
### Goroutines
### Channels


# Standard Library

## http/net

## sql


