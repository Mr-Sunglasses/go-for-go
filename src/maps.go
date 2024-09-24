package main

import "fmt"

func main()  {
	
	menu := map[string]float64{
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping through map
	for k,v := range menu {
		fmt.Printf("Key: %s and Value %f\n", k, v)
	}

	// ints as key type
	phoneBook := map[int]string{
		123: "mario",
		234: "luiji",
	}

	fmt.Println(phoneBook)
	fmt.Println(phoneBook[123])

	phoneBook[234] = "bowser"

	fmt.Println(phoneBook)
}