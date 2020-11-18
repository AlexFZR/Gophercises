package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Quit
//
//  Create a program that quits when a user types the
//  same word twice.
//
//
// RESTRICTION
//
//  The program should work case insensitive.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//
//   hey
//   HEY
//   TWICE!
//
//  go run main.go
//
//   hey
//   hi
//   HEY
//   TWICE!
// ---------------------------------------------------------

func main() {
	in := bufio.NewScanner(os.Stdin)

	words := make(map[string]bool)
	for in.Scan() {
		w := strings.ToLower(in.Text())
		if words[w] {
			fmt.Println("TWICE!")
			return
		}
		words[in.Text()] = true
	}
}
