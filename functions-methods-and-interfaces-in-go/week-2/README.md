# Notes

## First class values

- Treat functions as any other data type.
    - Variables can be declared with a function type.
    - Can be created dynamically.
    - Pass as arguments and/r returned as values.
    - Stored in a data structure.

``` go
var funcVar func(int) int

func incFn(x int) int {
    return x + 1
}

func main() {
    funcVar = incFn
    fmt.Print(funcVar(1)) //Same as incFn(1)
}
```

## Anonymous functions

- Don't need to be named. Useful when it's gonna be passed as argument (sort of lambda).

``` go
func applyIt(afunc func (int) int, val int) int {
    return afunc(val)
}

func main() {
    v := applyIt(func (x int) int {return x+1}, 2)
    fmt.Print(v)
}
```

## Functions as return values

- Closure = Function + Its environment.
    - When functions are passed/returned their environemnt comes with them. Careful when passing values that can be modified (e.g. references).

## Varadic functions

- Functions can take a variable number of arguments. Can be done with the ellipsis (`...`). The values are treated as a slice.

``` go
func MultipleVariables(vals ...int) int {
    total := 0
    // Can loop over the values
    for _, v := range vals {
        total = total + v
    }
    return total
}
```

- Can pass a slice to a variadic function

``` go
sli := []int{1, 2, 3}
MultipleVariables(sli...) // Note the ellipsis
```

## Deferred functions

- Call can be deferred ("delayed") until the surrounding function completes.
- Typically used for cleanup activities.

``` go
func main() {
    defer fmt.Println("Bye!")
    fmt.Println("Hello!")
}
```

- In the above example, the firt line would be executed once the `main()` function is done.
- Arguments are **evaluated immediately**, but the call is done later.