package main

import "fmt"

func main(){
	// x := 0

	// for x < 10 {
	// 	fmt.Println(x)
	// 	x++
	// }

	// for i := 0; i<5; i++{
	// 	fmt.Printf("%d\n",i)
	// }

	names := []string{"Kanishk", "Ayansh", "Saksham", "Shreya", "Aarav", "Ayushi", "Anika"}

	// for i := 0; i<len(names); i++{
	// 	fmt.Println(names[i])
	// }

	// for index, value := range names {
	// 	fmt.Println(index, value)
	// }

	// use _ if we don't want to use something
	// for _, value := range names {
	// 	fmt.Println(value)
	// }


	// value is the local copy of the value of index of names
	for index, value := range names {
		fmt.Println(index, value)
		value = "NewValue"
		fmt.Println(index, value)
	}

	fmt.Println(names)
}