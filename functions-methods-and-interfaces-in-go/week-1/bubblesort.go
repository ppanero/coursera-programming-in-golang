package main

import (
	"fmt"
	"strconv"
)

/*
 * Write a Bubble Sort program in Go. The program should prompt the user to
 * type in a sequence of up to 10 integers. The program should print the
 * integers out on one line, in sorted order, from least to greatest. Use
 * your favorite search tool to find a description of how the bubble sort
 * algorithm works.
 *
 * As part of this program, you should write a function called BubbleSort()
 * which takes a slice of integers as an argument and returns nothing. The
 * BubbleSort() function should modify the slice so that the elements are
 * in sorted order.
 *
 * A recurring operation in the bubble sort algorithm is the Swap operation
 * which swaps the position of two adjacent elements in the slice. You should
 * write a Swap() function which performs this operation. Your Swap() function
 * should take two arguments, a slice of integers and an index value i which
 * indicates a position in the slice. The Swap() function should return nothing,
 * but it should swap the contents of the slice in position i with the contents
 * in position i+1.
 *
 * The program should be tested with list of integers including negative numbers.
 */

func Swap(sli []int, position int) {
	aux := sli[position]
	sli[position] = sli[position+1]
	sli[position+1] = aux
}

func BubbleSort(sliToOrder []int) {
	swapped := false

	for i := 0; i < len(sliToOrder)-1; i++ {
		if sliToOrder[i] > sliToOrder[i+1] {
			Swap(sliToOrder, i)
			swapped = true
		}
	}

	if swapped {
		BubbleSort(sliToOrder)
	}
}

func main() {
	length := 10
	numbers := make([]int, length)

	// Read 10 integers
	fmt.Print("List to order: ")
	input := ""
	fmt.Scan(&input)
	for i := 0; i < length; i++ {
		num, _ := strconv.Atoi(input)
		numbers[i] = num
		fmt.Scan(&input)
	}

	BubbleSort(numbers)
	// Print sorted list
	fmt.Printf("%v\n", numbers)
}
