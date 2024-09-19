package main

import (
	"fmt"
	"sort"
)

func main(){

	// strings stdlib package
	// greetings := "Hello there my friends"

	// fmt.Println(strings.Contains(greetings, "Hello"))
	// fmt.Println(strings.ReplaceAll(greetings, "Hello", "hi"))
	// fmt.Println(strings.ToUpper(greetings))
	// fmt.Printf("%d\n", strings.Index(greetings, "ll"))
	// fmt.Println(strings.Split(greetings, " "))
	// fmt.Println(greetings)

	ages := []int{45, 55, 65, 75, 30, 20, 10, 11}

	sort.Ints(ages)
	fmt.Println(ages)

	index_of := sort.SearchInts(ages, 30)
	fmt.Println(index_of)

	names := []string{"Kanishk", "Ayansh", "Saksham", "Shreya", "Aarav", "Ayushi", "Anika"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "Ayushi"))
}