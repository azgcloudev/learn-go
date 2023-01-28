package main

import "fmt"

func main() {
	first := 100
	var second *int = &first

	first++

	fmt.Println("First:", first)
	fmt.Println("Second:", second)  // prints the pointer memory location
	fmt.Println("Second:", *second) // prints the pointer memory location value

	first1 := 100
	second1 := &first1

	first1++
	*second1++

	fmt.Println(first1)
	fmt.Println(*second1)

	// uninitialized pointers
	fmt.Println("\n--> uninitialized pointers <--")
	first2 := 100
	var second2 *int
	fmt.Println(second2)
	// fmt.Println(*second2) this creates an error if a pointer is used but not assigned a value
	// 	panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x481237]
	second2 = &first2
	fmt.Println(*second2)


	// pointing at pointers
	fmt.Println("\n--> pointing at pointers <--")
	first3 := 100
	second3 := &first3
	third3 := &second3

	fmt.Println(**third3)

}
