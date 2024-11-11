package main

import "fmt"

// must be outside of main function so it can be used in functions
var score = 99.5

// cannot use shorthand outside of functions
// scoreTwo := 50

func main() {
	sayHello("mario")
	showScore()

	for _, v := range points {
		fmt.Println(v)
	}
}
