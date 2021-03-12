# Notes

## Classes and Encapsulation

- Class: collection of data fields and methods (functions) that share a well-defined responsability.
- Object: instantiation of a class, a ser of values "inserted" into the template ("class").
- Encapsulation: Data can only be accessed/modified/etc. using specified methods. Helps to keep consistency, apply restrictions, etc.

## In Golang

- No *class* keyword.
- Methods are associated to data through "receiver" types, using a dor notation to call it.
- Cannot associate methods to built-in types, they need to be associated to types declared in the same package.
- By associating to a receiver type, it specifies an implicit method argument.
- Usually the receiver type is a struct, to allow for consistent data entities to be passed along implicitly.
- Functions to allow access to hidden data (from the package) start by capital letter, and therefore will be exported (usable from other package). Like `PrintDist`.
- A similar thing to "constructors" could be a function that initializes data of a receiver type. Same could apply for set/get/print.

```go
type Point struct {
    x float64
    y float64
}

func (p Point) DistToOrig() Point {
    t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
    return math.Sqrt(t)
}

func PrintDist() {
    p1 := Point(3, 4)
    fmt.Println(p1.DistToOrig())
}
```

## Point Receivers

- Method cannot modify the data inside the receiver, unless is passed by reference (otherwise it receives a copy).
- If the receiver object is large, it will copy a lot of data to the stack.
- No need to dereference a point receiver, no need to `*p`.
- No need to reference when calling a function with a point receiver, no need to `&p.Scale()`.
- Common good practice: Use all pointer receiver or none, do not missmatch methods. To avoid messing with (de)reference.
 
```go
func (p *Point) Scale() {
    p.x = p.x * 2
    p.y = p.y * 2
}

func PrintDist() {
    p1 := Point(3, 4)
    p1.Scale()
    fmt.Println(p1)
}
```