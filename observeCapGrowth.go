package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Observe the capacity growth
//
//  Write a program that appends elements to a slice
//  10 million times in a loop. Observe how the capacity of
//  the slice changes.
//
//
// STEPS
//
//  1. Create a nil slice
//
//  2. Loop 10e6 times
//
//  3. On each iteration: Append an element to the slice
//
//  4. Print the length and capacity of the slice "only"
//     when its capacity changes.
//
//  BONUS: Print also the growth rate of the capacity.
//
//
// EXPECTED OUTPUT
//
//  len:0               cap:0               growth:NaN
//  len:1               cap:1               growth:+Inf
//  len:2               cap:2               growth:2.00
//  ... and so on.
//
// ---------------------------------------------------------

func main() {
	slice, oldCap := []int{}, 1.

	for len(slice) < 10e6 {

		capacity := float64(cap(slice))
		if capacity ==0 || capacity != oldCap {
			fmt.Printf("len:%-15d cap:%-15g growth:%.2f\n", len(slice), capacity, capacity/oldCap)
		}
		oldCap = capacity
		slice = append(slice, 1)
	}
}
