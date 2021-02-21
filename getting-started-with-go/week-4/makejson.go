package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

/*
 * Write a program which prompts the user to first enter a name, and then enter an address.
 * Your program should create a map and add the name and address to the map using the keys
 * “name” and “address”, respectively. Your program should use Marshal() to create a JSON
 * object from the map, and then your program should print the JSON object.
 */
func main() {

	var name, addr string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Name: ")
	scanner.Scan()
	name = scanner.Text()
	fmt.Printf("Address: ")
	scanner.Scan()
	addr = scanner.Text()

	var person = map[string]string{"name": name, "address": addr}

	var pjson, _ = json.Marshal(person)
	fmt.Printf("%q", pjson)
}
