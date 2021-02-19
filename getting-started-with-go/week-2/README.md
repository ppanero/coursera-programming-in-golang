# Notes

## Pointers

- A pointer is an address to data in memory.
- `&` returns the address.
- `*` returns the data at the address (value), called dereferencing (e.g. `*x` would return the value at the address `x`, which assume `x`'s value is a memory address).
- `new()` creates a variable and returns a **pointer** to the variable. To access its value has to be accessed with `*`.

## Scopes

- Available in the block (`{}`) and the ones below in the hierarchy (_inner blocks_).
- Implicit blocks:
    - Universe block - all Go source
    - Package block - all source in a package
    - File block - all source in a file (e.g. `.go`)
    - `if`, `for`, `switch`, etc.
- Lexical scoping

## Memory deallocation

- Stack: function scope, deallocation is _automatic_ when the function completes.
- Heap: Universe, package, file scope, is _persistent_.
    - Manually (e.g. `free()`): error-prone but fast.
    - Garbage collection

## Garbage collection

- Hard to figure out when to deallocate a variable (e.g. "Are there pointers to the mem address?").

```go
func foo() *int {
    x := 1
    return &x
}

func main() {
    var y *int
    y = foo()
    fmt.Printf("%d", *y)
}
```

- Usually done by interpreters (e.g. JVM). But Go **enables it**!
- As a programmer there is no need to know if memory should be allocated in the heap or the stack.
- There is a performance hit, but it trades off by making programmers life easier.

## Integers, Floats, Strings, Constants

- Type conversion:
    - Not always possible
    - Done using the `<Type>()` operation (e.g. `int32(x)`).

**Integers**

- Many types: _int8, int16, int32, int64_ and the same for _uint_ (unsigned, can allocate a larger value).
- Operators for arithmetic (`+ - * / % << >>`), comparison (`== != > < >= <=`), boolean (`&& ||`) and bit-wise operations (`& |`).

**Floating point**

- Precision errors are usually hard to catch. As a rule of thumb go larger, even though it can hit on performance. `float32` ~6 digit precision, `float64` ~15 digits precision, etc.
- `complex(real, imaginary)` numbers

**Characters**

- ASCII: Each character ir represented by an 8-bit code. Max 256 characters (228 in reality, one bit is used for other purposes).
- Unicode: 32-bit code, better when dealing with not-only-english text.
- **UTF-8 (Go default)**, starts by 8-bit code but can extend up to 32 and the first 228 match the ASCII codes.
- Code point = Unicode character; _Rune_ = a code point in Go.

**Strings**

- Strings are read-only. Cannot modify an existing string, can create a new one with a modified value.
- Sequence of arbitrary bytes, where each byte is a _rune_.
- Useful functions: `IsDigit(r rune)`, `IsSpace(r rune)`, `ToUpper(r rune)`, `ToLower(r rune)`.
- _Strings_ package: compare, search, contain, replace, trim spaces, etc. functions over a _whole string_ not only over a single rune.
- _Strconv_ package: conversion operations, e.g. `Atoi`, `Itoa`.

**Constants**

- Can be declared with `const x int`
- If the variable is one-hot (e.g. days of the week), then they can be created with `iota`. The caveat is that the values will be different (e.g. _monday_ different from _tuesday_) but the actual value might be something different (e.g. _250_ instead of _monday_). Useful for enumerations. e.g.:

``` go

type Grades int
const (
    A Grades = iota
    B
    C
    D
    F
)
```

- Note: There is no need to specify the type in all the constants of the enumeration, just the first one. Iota will internally assign an integer starting by `1` (but the programmer should not rely on that).

## Control flow

- `if/else` statements.
- `for` loops. `break` and `continue`, although they are not recommended.
- `switch/case` statements. There is no need for `break` statements, the case breaks automatically.
- _tagless_ `switch`, instead of comparing a value to constants it will evaluate expressions. e.g.

``` go
switch {
    case x >1:
        ...
    case x < -1:
        ...
    default:
        ...
}
```

## User input

- Using the `scan` function.
- Takes a pointer as an argument.
- Typed data is written to a pointer.
- Returns the number of scanned items and an error (if happened).
