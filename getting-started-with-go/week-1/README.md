# Notes

## Why Go?

- Runs fast
    - Compiled language (no run-time translations)
- Garbage collection.
    - Automatic memory management (allocation and deallocation).
    - Usually a feature of interpreted languages.
- Simpler objects. Easier to code, less complex features.
- Concurrency is efficient. It was created with concurrency in mind.

## Go Tool

- `go build [pacages or list of .go files]` compiles the program.
- `go doc` prints the documentation for a package.
- `go fmt` formats source code files.
- `go get` downloads packages and installs them.
- `go list` lists all installed packages.
- `go run` compiles and executes the compiled file.
- `go test` runs the `*_test.go` files.

## Coding in Go

- Download and install from [here](https://golang.org).
- Standard organization of workspace/files:
    - src/ (source code)
    - pkg/ (packages, libraries)
    - bin/ (executables)
- Workspace dir is defined by `GOPATH` env var.
- Go tools assume the code is in the `GOPATH`.
- Packages:
    - First line of every file names the package (e.g. `package usefulpkg`).
    - The *must* be a `main` package, where execution starts.
    - Compiling the main package generates an executable file/program.
    - `import`s are resolved from `GOROOT` and `GOPATH` env vars directories.

## Object Oriented

- Less features, it is weakly OO.
- Object ~ User-Defined type (set of related attrs and functions).
- Does **not use `class`**, it uses `structs` with associated methods.
    - No inheritance.
    - No constructors.
    - No generics.

## Concurrency

- Might not run at the same time, but are alive at the same time (i.e. concurrent, maybe not parallel if for example executed in single core).
- Management of task execution (`Goroutines`)
- Communication between tasks (`Channels`)
- Synchronization between tasks (`Select`)  

## Variables

- Names
    - Must start by a letter.
    - Accept any amount of letters, digits and underscores.
    - Are case sensitive.
    - There are reserved keywords (`if`, `case`, `package`).
- Variables are typed and must be declared, e.g. `var x int`.
    - Can declared multiple of the same type in one line, e.g. `var x, y int`.
- Variables have to be initialized:
    - `var x int = 100`, specified type (preferred)
    - `var x = 100`, type is inferred by the compiler.
    - `var x int\n x = 100`.
    - If not initialized variables are initialized to the default "empty" value for a type, e.g. `0` for `int`, `""` for `string`.
    - `x := 100` performs a declaration and initialization in the same line. The type is set to the type of the right hand side value (inferred). Only possible inside a function.
- Basic types:
    - Integers
    - Floating points (Decimal). Might use different hardware below. Use integer when possible.
    - Strings. Represented in unicode.
- Can define type *aliases*, e.g. `type Celsius float64`, which is the same type (same functions, etc.) but will improve clarity in the code. Then you can use it, e.g. `var temp Celsius`.


