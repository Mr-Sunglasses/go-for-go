package main

import "fmt"

/*
Non-Pointer Value
String Ints floats booleans arrays structs
*/

func updateValue(s string) {
	// this will not work
	s = "update"
	fmt.Printf("Update Value %q inside a function\n",s)
}

func updateValueThatWorks(s string) string  {
	s = "updated"
	return s
}

func addItemToMap(s map [int]string)  {
	// This will work as they are pointer wrapper values
	s[120] = "Apple"
}

/* 
Pointer Wrapper Value
Slices
Maps
function
*/
func main() {
	v := "noupdate"
	updateValue(v)
	v = updateValueThatWorks(v)
	fmt.Println(v)

	vi := map[int]string{
		123: "JLO",
		234: "JTO",
	}

	addItemToMap(vi)

	fmt.Printf("%v\n", vi)
}
