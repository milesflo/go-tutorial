package main

import "fmt"

func functionNam() {
	//
}

func calculateBill(price int, no int) int { // Parameters may be typed, as well as the return type
	var totalPrice = price * no
	return totalPrice
}

func rectProps(length, width float64)(float64, float64) { // Those two float64s represet the return values
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter // Two values can be returned from a single function
}

func main() {
	price, no := 90, 6
	totalPrice := calculateBill(price, no)

	area, _ := rectProps(10.8, 5.6) // You can include a _ blank identifier to drop a value

	fmt.Println(totalPrice)
	fmt.Println(area)
}