package main

import (
	"fmt"
	// "math/rand"
)

func main() {
	fmt.Println("Hello, Go")
	fmt.Println(20 + 20)
	fmt.Println(20 + 30)
	fmt.Println("Hello\n")
	fmt.Println(`Hello\n`)

	// constants
	const price float32 = 275.00
	const tax float32 = 27.50
	fmt.Println(price + tax)

	// untype const (makes go to automatic convert it)
	const quantity = 2                          // int
	fmt.Println("Total:", quantity*(price+tax)) // quantity becames a float32

	// IOTA
	// continue value aassigment starting from 0
	const (
		Watersports = iota
		Soccer
		Chess
	)
	fmt.Println("Values are:", Watersports, Soccer, Chess)

	const price1, tax1 float32 = 275, 27.50
	const quantity1, inStock1 = 2, true
	fmt.Println("Total:", quantity1*(price1+tax1))
	fmt.Println("In stock: ", inStock1)
}
