package main

import "fmt"

func main(){
	// Arrays: fixed length
	var ages[3]int = [3]int{1,2,3}
	ages[1] = 20
	var myAges = [5]int{30,40,50,60,70}
	myNames := [2]string{"Kanishk", "Kayansh"}
	fmt.Printf("%v \n", ages)
	fmt.Println(myAges, len(myAges))
	fmt.Printf("%T", len(myNames))

	// slices( they are same as arrays and even use arrays under the hood but they can be of any length and can be appendable)
	var scores = []int{100,50, 60}
	scores[2] = 25
	allScoreNames := []string{"One"}
	allScoreNames = append(allScoreNames, "Two")
	fmt.Println(allScoreNames, len(allScoreNames))

	// slices range
	rangeOne := myNames[1:2] // Include position 1 but not positon 2 
	rangeTwo := myAges[2:] // Include positon 2 and go till the end
	rangeThree := myAges[:len(myAges)-1] // Include All before the postion but not include the last, ex [:4] include 0,1,2,3 but not 4
	fmt.Printf("%v\n%v\n%v\n",rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "Kooper")
	fmt.Println(rangeOne)
}