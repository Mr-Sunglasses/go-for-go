package main

import "fmt"

func UpdateValue(s *string)  {
	*s = "Kanishk"
}

func main()  {
	name := "Yuvraj"

	m := &name

	fmt.Println(m)
	fmt.Println(&m)
	fmt.Println(&name)
	fmt.Println(*m)ˀ

	// UpdateValue(&name)

	// fmt.Println(name)

	// fmt.Println(&name)

	// m := &name

	// *m = "Vivo"

	// fmt.Printf("value of name is: %s and value of m is: %s\n", name, *m)
}