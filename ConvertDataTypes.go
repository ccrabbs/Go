package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Convert Data Types
	fmt.Println("Convert Data Types")
	myInteger := "42"
	numInt, err := strconv.Atoi(myInteger)
	fmt.Printf("num:%v %T %v\n", numInt, numInt, err)
	myFloat := "123.99999"
	numFloat, err := strconv.ParseFloat(myFloat, 64)
	fmt.Printf("num:%v %T %v\n", numFloat, numFloat, err)
	numDecimal := float64(numInt)
	fmt.Printf("Decimal %.2f %T\n", numDecimal, numDecimal)
	var char byte = 'A'
	var str = string(char)
	fmt.Printf("Char %v String %v\n", char, str)
}
