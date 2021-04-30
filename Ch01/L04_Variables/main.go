package main

import "fmt"

func main()  {
	var quantity int
	var length, width float64
	var customerName string
	shortVarDec := true

	quantity = 4
	length, width = 1.2, 2.4
	customerName = "Damon Cole"

	var myInt int

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with area of", width * length, "square meters")
	fmt.Println(myInt)
	fmt.Println(shortVarDec)
}
