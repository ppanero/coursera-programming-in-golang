# Notes

## Functions

**Why?**
- Reusability.
- Clarity.
- Encapsulation/Abstraction.

**How?**
- Declaration: `func CamelCaseNamedFunction (arg type) returnType `

```go
func SumTwo (arg int, arg2 int) int {
    return arg + arg2
}

func main() {
    fmt.Printf("%d", SumTwo(1, 2))
}
```

- If the type or arguments is the same they can be comma separated (e.g. `func foo (x, y int)`).
- There can be more than one return value (e.g. `func foo2 (x, y int) (int, string)`).

## Arguments passing

**Call by value**
- A copy of the argument is passed.
- Whatever happens in the function **has no effect** on the variable outside of it.
- Advantage: Data encapsulation. Cannot affect the caller, encapsulating the propagation of errors.
- Disadvantage: Copying time. If objects are large it will take a long time.

**Call by reference**
- Also called "pass by pointer" or "pass poiter". Although it is not 100% correct to say so, see more [here](https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go).
- A pointer to the variable memory address is passed. The function declaration will have `*type` as argument type, it's value will be accessed with `*argName`. Usually no return is needed, since the value was already modified.
- Whatever happens in the function **affects** the variable outside of it.
- Advantage: Copying time. It is fast even for large objects, since it will copy only a pointer.
- Disadvantage: Data encapsulation. Bugs/errors will spread.

```go
func SumOne (arg *int) {
    *arg = *arg + 1
}

func main() {
    x := 1
    SumOne(&x)
    fmt.Printf("%d", x)
}
```

**Passing arrays and slices**
- Do it by reference to avoid copying memory problems.
- Pass slices instead of arrays.
    - Slices contain a pointer to the array.
    - Passing a slice copies the pointer to the array (along with length and capacity).
    - Therefore it avoids handling the pointers.

``` go
func foo(sli []int) int {
    sli[0] = sli[0] + 1
}

func main() {
    a := []int{1, 2, 3}
    foo(a) //Note no umpersand (&)
    fmt.Print(a)
}
```

## Well written functions and guidelines

- Understandable, clear (e.g. named correctly).
- Where and how data is used (e.g. how could it have gotten incorrect? was it passed by reference?).
- Perform one operation only. If more than one are needed, split in several functions and create another one to wrap them up.
- The less parameters the better. If the function requires a lot of arguments, it might be a sign of bad functional cohesion. It can also be due to bad/lack of structure definition.

**Debugging principles**
- Function is written incorrectly. --> Functions need to be understandable.
- Data that the function uses is incorrect (e.g. arguments, file input). --> Data needs to be tracable.
