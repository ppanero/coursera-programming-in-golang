package main

import (
	"fmt"
	"math"
)

/*
 * Let us assume the following formula for displacement s as a function of
 * time t, acceleration a, initial velocity vo, and initial displacement so.
 *
 *	 s = 1/2 * a * t^2 + v_0 * t + s_0
 *
 * Write a program which first prompts the user to enter values for
 * acceleration, initial velocity, and initial displacement.
 * Then the program should prompt the user to enter a value for time and the
 * program should compute the displacement after the entered time.
 *
 * You will need to define and use a function called GenDisplaceFn() which
 * takes three float64 arguments, acceleration a, initial velocity v_0, and
 * initial displacement s_0.
 * GenDisplaceFn() should return a function which computes displacement as
 * a function of time, assuming the given values acceleration, initial
 * velocity, and initial displacement. The function returned by
 * GenDisplaceFn() should take one float64 argument t, representing time,
 * and return one float64 argument which is the displacement travelled
 * after time t.
 *
 * For example, let’s say that I want to assume the following values for
 * acceleration, initial velocity, and initial displacement:
 * a = 10, vo = 2, so = 1. I can use the
 * following statement to call GenDisplaceFn() to generate a function fn
 * which will compute displacement as a function of time.
 *
 * fn := GenDisplaceFn(10, 2, 1)
 *
 * Then I can use the following statement to print the displacement after 3
 * seconds.
 *
 * fmt.Println(fn(3))
 *
 * And I can use the following statement to print the displacement after 5
 * seconds.
 *
 * fmt.Println(fn(5))
 */

func RequestValues() (float64, float64, float64, float64) {
	var acc, v0, s0, time float64

	fmt.Print("Acceleration: ")
	fmt.Scanln(&acc)
	fmt.Print("Initial velocity: ")
	fmt.Scanln(&v0)
	fmt.Print("Initial displacement: ")
	fmt.Scanln(&s0)
	fmt.Print("Time: ")
	fmt.Scanln(&time)

	return acc, v0, s0, time
}

func GenDisplaceFn(acc, v0, s0 float64) func(time float64) float64 {
	return func(time float64) float64 {
		//  s = 1/2 * a * t^2 + v_0 * t + s_0
		partOne := 0.5 * acc * math.Pow(time, 2)
		partTwo := v0 * time
		return partOne + partTwo + s0
	}
}

func main() {
	acc, v0, s0, time := RequestValues()
	displacementFn := GenDisplaceFn(acc, v0, s0)
	fmt.Printf("%f", displacementFn(time))
}
