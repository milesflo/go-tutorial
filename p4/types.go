package main

import "fmt"

func main() {
	a := true // Bools are lowercase in Go
	b := false
	fmt.Println("a:",a,"b:",b)
	c := a && b // AND operand
	fmt.Println("c:", c)
	d := a || b // OR operand
	fmt.Println("d:", d)

	e, f := 5.67, 8.97
	fmt.Printf("type of a %T b %T\n", e, f) // %T will select type of variable
	sum	:= e + f
	diff := e - f
	fmt.Println("sum", sum, "diff", diff) // Float point math error

	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2) // No error since they're int type

	bigint := 6 + 7i // complex number
	fmt.Println(bigint)

	

}
