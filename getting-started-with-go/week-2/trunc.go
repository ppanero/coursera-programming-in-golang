package main

import "fmt"

/*
 * The program has to print the truncated part of a floating point number.
 * For example 1 for 1.1, 39 for 39.89.
 * It has to run two times.
 */
func main() {

	for i := 0; i < 2; i++ {
		// Request number to the user
		var float_num float64
		fmt.Scan(&float_num)
		// Print truncated number
		fmt.Printf("%d\n", int(float_num))
	}
}
