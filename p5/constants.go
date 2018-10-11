// Constants are unique in that they must be known at compile time


package main

import (
	"fmt"
)

func main() {
	const a string = "I love Go!"
	
	fmt.Println(a)
	
	var defaultName = "Sam"
	type myString string // Create a type `myString` which is a copy of string
	var customName myString = "Sam"
	// customName = defaultName // This will throw an error since they're not of the same datatype 
}