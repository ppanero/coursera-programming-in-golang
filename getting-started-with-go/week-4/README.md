# Notes

## RFCs (Request For Comments)

- Definitions of protocols, formats, standards, etc. (e.g. HTML, HTTP, URI).
- Golang has packages with implementations of the most important RFCs (e.g. "net/http").

## JSON (JavaScript Object Notation)

- Started in JavaScript, spread widely.
- RFC 7159.
- Structured information in KV pairs.
- Unicode.
- Smallest (compact) while still human readable.
- Marshalling (object-to-json) can be done with the _json_ package (e.g. `json.Marshal(struct1)`). It returns a byte array withe the json representation and an error if it happened. The opposite operation can be done with the `json.Unmarshal(bytearr, &struct1`.

## File access

- Linear access, not random access.
- Basic operations: open, read, write, close and seek.

**ioutil**

- Basic functions:
    - `byte_arr, err := ioutil.ReadFile(<path_to_file>)`
    - No need for open/close statements, it is done automatically.
    - Large files can cause performance problems and eventual problems. It reads all into memory (so it is limited by the RAM of the computer).
    - `err := ioutil.WriteFile(<path_to_file>, byte_arr, <permission>)`, permissions are unix style (e.g. 0777).
    - No append possible, it erases the content upon opening.

**os**

- Gives more flexibility.
- Open files with `os.Open()`, returns a file descriptor. Or create with `os.Create()`.
- Close files with `os.Close()`.
- Read a file into a byte array (controlled amount) with `os.Read()`.
- Write to a file with `os.Write()`. Can append.
