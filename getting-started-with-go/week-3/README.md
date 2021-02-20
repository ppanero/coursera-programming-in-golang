# Notes

## Arrays

- **Fixed length** series of elements of a chosen type.
- 0 indexed
- All elements initialized to the _zero_ value and access using the subscript notation `[ ]`.

``` go
var x [5]int
x[0] = 2
fmt.Printf(x[1]) // Would print 0
```

- An _array literal_ is an array with predefined values (e.g. `var x [3]int = [3]{1, 2, 3}`).
    - `...` for size in array literal infers size from number of initializers (e.g. `x := [...]int{1, 2, 3}`).
- Use a `for` loop with `range` to iterate through the array. `range` returns the index and the value.

## Slices

- A "window" on an underlying array.
- **Variable size**, up to the whole array.
- Properties:
    - _Pointer_ indicates the start of the slice.
    - _Length_ is the number of elements in the slice. Can be obtained with `len()`
    - _Capacity_ is the max number of elements (From start till end of array). Can be obtained with `cap()`.
- Slices include the first and up to (but not including) the last (e.g. `arr[1:3]` would include elements 1 and 2, but not 3).
- Slices can overlap. Therefore, if you change one you are changing the other, since they refer to the same underlying array.
- A _slice literal_ can be used to initialize a slice. It creates the underlying array and a slice referencing it. The slice will be the full array (i.e. length == capacity). e.g. `sli := []int{1, 2, 3}`. Note the difference is the lack of `...`.

## Variable slices

- `make(type, length, [capacity])` can create slices.
    - If capacity is not specified, the length will be the capacity.
    - Type has to be set to array e.g. `[]int`.
- Useful functions:
    - `append(slice, elem)`, if needed it will increase the **length of the slice** (creates a new larger slice). If it reaches the max of the underlying array, it creates a new larger one (there is a time penalty for that).

## Hash tables

- Stores key/value pairs. Keys are unique.
- Uses a _hash_ function to compute the slot for the key (where to insert the value).
- Access using the subscript notation with the `key` (e.g. `hash["joe"]`)
- Faster lookup. Constant time (O(1)), an array is linear time (O(n)).
- May have collisions if the keys hash to the same slot. They are very rare, and Go handles them automatically, however they have a time penalty.

### Maps

- An implementation of a hash table, `var exMap map[key_type]value_type`.
- Can be created using `make()`.
- Can be created with a _map literal_.
- Access returns _zero_ if not present.
- Two-value assignments tests for existence of the key `id, p := exMap["joe"]` where `p` will be `true` if the key `joe` is on the map.
- To itereate through a map use a `for` loop with `two-value assignemnt range`. Each iteration will return a `key, value` pair.

```go
for key, val := range exMap {
    fmt.Println(key, val)
}
```

## Structs

- Aggregate data type.
- Groups together other objects of arbitrary type (e.g. a _Person_ is composed of _Name, Address, Phone_).
- Fields are accessed using the _dot notation_

```go
type Person struct {
    name string
    addr string
    phone string
}

p1 := new(Person) //Initalize a new person to zero values
p2 := Person{name: "Joe", addr: "street", phone: "123"} //Initialized struct literal
x = p1.addr //Access example
```

