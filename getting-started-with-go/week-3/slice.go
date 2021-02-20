package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
 * Write a program which prompts the user to enter integers and stores the
 * integers in a sorted slice. The program should be written as a loop.
 * Before entering the loop, the program should create an empty integer slice
 * of size (length) 3. During each pass through the loop, the program prompts
 * the user to enter an integer to be added to the slice. The program adds the
 * integer to the slice, sorts the slice, and prints the contents of the slice
 * in sorted order. The slice must grow in size to accommodate any number of
 * integers which the user decides to enter. The program should only quit
 * (exiting the loop) when the user enters the character ‘X’ instead of an
 * integer.
 */
func main() {

	var input string
	sli := make([]int, 0, 3) // length 3 as specified in the task description

	fmt.Scan(&input) // Assumes input is int or X, no input checkup is performed
	for input != "X" {
		i, _ := strconv.Atoi(input)
		sli = append(sli, i)
		// Sort each time. To make it more efficient use
		// binary search + sorted insert
		sort.Slice(sli, func(i, j int) bool { return sli[i] < sli[j] })
		fmt.Printf("%v\n", sli)
		fmt.Scan(&input)
	}
}
