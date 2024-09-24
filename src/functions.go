package main

import (
	"fmt"
	"math"
)

func sayGreetings(n string) string  {
	return fmt.Sprintf("Good morning %s\n", n)
}

func sayBye(n string)  {
	fmt.Printf("Good Bye %s\n", n)
}

func cycleName(n []string, f func(string) string){
	for _, value := range n{
		fmt.Printf("%s", f(value))
	}
}

func circleArea(r float64) float64{
	return math.Pi * r * r
}

func main(){
	// sayGreetings("mario")
	// sayGreetings("luigi")
	// sayBye("mario")

	cycleName([]string{"cloud", "tifa", "barret"}, sayGreetings)
	cycleName([]string{"cloud", "tifa", "barret"}, sayBye)

	areaOne := circleArea(10.5)

	fmt.Printf("%0.2f\n", areaOne)
}