package main

// ---------------------------------------------------------
// EXERCISE: Infinite Kill
//
//  1. Create an infinite loop
//  2. On each step of the loop print a random character.
//  3. And, sleep for 250 milliseconds.
//  4. Run the program and wait a couple of seconds
//     then kill it using CTRL+C keys
//
// RESTRICTIONS
//  1. Print one of those characters randomly: \ / | -
//  2. Before printing a character print also this
//     escape sequence: \r
//
//     Like this: "\r/", or this: "\r|", and so on...
//
//  3. NOTE : If you're using Go Playground, use "\f" instead of "\r"
//
// EXPECTED OUTPUT
//  - The program should display the following messages in any order.
//  - And, the first character (\, -, /, or |) should be randomly
//    displayed.
//  - \r or \f sequence will clear the line.
//
//  \ Please Wait. Processing....
//  - Please Wait. Processing....
//  / Please Wait. Processing....
//  | Please Wait. Processing....
// ---------------------------------------------------------

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {



	for {
		var c string
		switch rand.Intn(4) {
		case 0:
			c = "\\"
		case 1:
			c = "-"
		case 2:
			c = "/"
		case 3:
			c = "|"
		}
		fmt.Printf("\r%s Please Wait. Processing....", c)
		time.Sleep(time.Millisecond * 250)
	}

}
