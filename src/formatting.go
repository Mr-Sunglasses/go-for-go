package main

import "fmt"

func main(){
	age := 35
	name := "shauna"

	// Print
	fmt.Print("hello, ")
	fmt.Print("World! \n")
	fmt.Print("In the new line!")

	// Println
	fmt.Println("hello ninjas!")
	fmt.Println("GoodBye Ninjas!")
	fmt.Println("my age is:", age, "and my name is:", name)

	// String Formatting
	/* 
		%v 
		%q
		%T
		%f
		%0.1f for one decimal point
		%0.2f for two decimal point
	*/
	fmt.Printf("My age is %T and my name is %T \n", age, name)

	// Sprintf (Return formatted string)

	mystring := fmt.Sprintf("Hello World, Java age is %v", age)
	fmt.Printf("%v \n", mystring)
}