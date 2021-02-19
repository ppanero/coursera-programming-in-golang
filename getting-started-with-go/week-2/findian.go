package main

import (
	"fmt"
	"regexp"
)

/*
 * The program has to print "Found!" when a string with the characters
 * 'i', 'a' and 'n' is entered, and "Not found!" if it does not contain them.
 * It is not specified what to do if only part of the characters are found
 * (e.g. "a" and "n" but not "i").
 * It has to run two times.
 *
 * For simplicity on the regexpresion, only i-a-n order is taken into account.
 */
func main() {

	r, _ := regexp.Compile("([a-zA-Z0-9]*i[a-zA-Z0-9]*a[a-zA-Z0-9]*n[a-zA-Z0-9]*)")

	for i := 0; i < 2; i++ {
		// Request string to the user
		var user_input string
		fmt.Scan(&user_input)
		// Print truncated number
		if r.MatchString(user_input) {
			fmt.Printf("Found!\n")
		} else {
			fmt.Printf("Not found!\n")
		}
	}
}
