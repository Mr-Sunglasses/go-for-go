package main

import "fmt"

var score float64 = 99.6

func main() {

	// var score float64 = 99.6 now score is not in package scope.

	sayGreetings("mario")

	for _, value := range points{
		fmt.Printf("%d\n", value)
	}

	showScore()
}