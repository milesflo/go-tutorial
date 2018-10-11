package main

import (
	"fmt"
	"math"
)

func main() {
	var age int = 29 // variable declaration with initial value
	fmt.Println("my age is", age)

	var newage = 54 // type will be inferred
	fmt.Println("my new new age is", newage)

	var width1, height1 int = 100, 50 // Declare mutiple variables of same type at once

	fmt.Println("width1 is", width1, "height1 is", height1) // String concatenation
	var width2, height2 = 100, 50 // Type not required

	fmt.Println("width2 is", width2, "height2 is", height2) // String concatenation

	var width3, height3 int // Init multiple with no default value

	fmt.Println("width3 is", width3, "height3 is", height3)
	width3 = 100
	height3 = 50
	fmt.Println("width3 is", width3, "height3 is", height3)

	var (
		name4 	= "Miles"
		age4	 	= 22
		height4	int
	)

	fmt.Println("my name is", name4, ", age is", age4, "and height is", height4)

	name5, age5 := "Miles", 22 // Short hand variable declaration

	fmt.Println("My name is", name5,"age is", age5)

	a, b := 20, 30 // declare variables a and b
	fmt.Println("a is", a, "b is", b)
	b, c := 40, 50 // b is already declared but c is new
	fmt.Println("b is", b, "c is", c)
	b, c = 80, 90 // assign new values to already declared variables b and c
	fmt.Println("changed b is", b, "c is", c)

	d,e := 123.6, 534.2
	min := math.Min(d,e)
	fmt.Println("The smallest of those two numbers was", min)
}

