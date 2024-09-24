package main

import (
	"fmt"
	"os"
)


type bill struct {
	name string
	items map[string]float64
	tip float64
}

func newBill(name string) bill {
	b := bill{
		name: name, 
		items: map[string]float64{}, 
		tip: 0,
	}
	return b
}

// Reciver Function
func (b *bill) format() string {
	fs := "Bill Breakdown: \n"
	var total float64 = 0
	
	//list items
	for key, value := range b.items {
		fs += fmt.Sprintf("%-25v ....$%0.2f \n", key+":", value)
		total += value
	}

	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "tip:", b.tip)
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "total:", total+b.tip)
	
	
	return fs
}

//Update tip
func (b *bill) Addtip(tip float64)  {
	// Derefering the ptr
	(*b).tip += tip
}

// Add items
func (b *bill) Additems(item string, price float64) {
	// (*b).items[item] = price 
	// We can also use it without derefrencing
	b.items[item] = price
}

// Save bill to file
func (b *bill) Savebill()  {
	data := []byte(b.format())
	path := fmt.Sprintf("bills/%s_bill.txt", b.name)

	err := os.WriteFile(path, data, 0644)

	if err!=nil{
		panic(err)
	}

	fmt.Println("Bill is saved to file")
}