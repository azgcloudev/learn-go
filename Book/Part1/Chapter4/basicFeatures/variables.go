package main

import "fmt"

func main() {

	// var price float32 = 275.00
	// var tax float32 = 27.50
	// fmt.Println(price + tax)
	// price = 300
	// fmt.Println(price + tax)

	//omitting data type
	var price = 275.00
	var price2 = price
	fmt.Println(price)
	fmt.Println(price2)

	// multiple variables with single statement
	var price1, tax1 = 275.00, 27.50
	fmt.Println(price1 + tax1)

	// short declaration syntax
	price3 := 275.00
	fmt.Println(price3)
	price4, tax4, inStock4 := 275.00, 27.50, true
	fmt.Println("Total: ", price4 + tax4)
	fmt.Println("In stock: ", inStock4)
}
