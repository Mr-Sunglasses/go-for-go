package main

import (
	"fmt"
	"strings"
)

func main()  {
	name := []string{"Kanishk", "Ayansh", "Ayushi", "Shreya", "Aarav", "Saksham"}
	ahe := 21
	fmt.Printf("%s\n", name)
	fmt.Println(ahe <= 50)
	fmt.Println(ahe >= 50)
	fmt.Println(ahe == 21)
	fmt.Println(ahe != 50)

	if ahe < 30 {
		fmt.Println("WOW you are young do something cool")
	}else if ahe < 40{
		fmt.Println("WoW still young")
	}else {
		fmt.Println("You are old! Time to Retire")
	}

	for index, value := range name {
		if strings.ToLower(value)  == "shreya"{
			fmt.Println(index)
			break
		}
	}

	for index, value := range name{
		if index == 1{
			fmt.Printf("Continue at index %d\n", index)
			continue
		}else if index == len(name) - 1{
			fmt.Printf("Breaking at Last Index: %d\n", index)
		}else{
			fmt.Printf("Index: %d\t Value: %s\n", index, value)
		}
	}
}