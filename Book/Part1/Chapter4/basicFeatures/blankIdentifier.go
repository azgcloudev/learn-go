package main

import "fmt"

func main() {

	var _ = "Alice" // not using the variable, provides no errors because of the "_" character black identifier
	fmt.Println("Blank identifier")
}
