package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Input (prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func Createbill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := Input("Please enter your name: ", reader)

	b := newBill(name)

	fmt.Println("Created the bill - ", b.name)

	return b
	
}

func Promptopts (b bill ){

	reader := bufio.NewReader(os.Stdin)

	opts, _ := Input("Choose option (a - to add item, s - to save bill, t - to add tip): ", reader)

	switch opts {
	case "a":
		name, _ := Input("Please Add the item name: ", reader)
		price, _ := Input("Please Enter the item price in ($)....: ", reader)


		p, err := strconv.ParseFloat(price, 64)
		if err !=nil {
			fmt.Println("The price must be a number")
			Promptopts(b)
		}

		b.Additems(name, p)

		fmt.Printf("Added the Item: %s and price: %0.2f to the bill \n", name, p)
		Promptopts(b)
	case "s":
		b.Savebill()
		fmt.Printf("Your Bill is saved in file: %q", b.name+"_bill.txt")
	case "t":
		tip, _ := Input("Please Enter the Amount of tip in ($)....:", reader)
		
		t, err := strconv.ParseFloat(tip, 64)
		if err !=nil {
			fmt.Println("The tip must be a number")
			Promptopts(b)
		}

		b.Addtip(t)

		fmt.Printf("The tip of %0.2f is added to the bill. \n", t)
		Promptopts(b)

	default:
		fmt.Println("That was not a valid option")
		Promptopts(b)
	}

}

func main()  {

	myBill := Createbill()
	Promptopts(myBill)



	// fmt.Println(Createbill())
	// mybill := newBill("Mario's Bill")
	// fmt.Println(mybill)
	
	// newVar := bill{
	// 	name: "Kanishk",
	// 	items: map[string]float64{},
	// 	tip: 0,
	// }

	// fmt.Println(newVar)

	// mybill.Updatetip(10)

	// mybill.Additems("Biriyani", 6.99)

	// fmt.Println(mybill.format())

	// mybill := Createbill()
	// fmt.Println(mybill)

}