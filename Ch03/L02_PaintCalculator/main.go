package main

import "fmt"

var metersPerLiter float64

func paintNeeded(width, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", height)
	}

	area := width * height
	return area / metersPerLiter, nil
}

func double(num *int) {
	*num *= 2
}

func createPointer() *float64 {
	myFloat := 98.5
	return &myFloat
}

func printBoolPointer(pointer *bool)  {
	fmt.Println(*pointer)
}

func main()  {
	metersPerLiter = 10.0
	var total float64

	amount, err := paintNeeded(4.2, -3.0)
	if err != nil {
		println(err.Error())
	} else {
		total += amount
		fmt.Printf("%.2f liters needed\n", amount)
	}

	amount, err = paintNeeded(5.2, 3.5)
	if err != nil {
		println(err.Error())
	} else {
		total += amount
		fmt.Printf("%.2f liters needed\n", amount)
	}

	fmt.Printf("Total: %0.2f liters\n", total)
	fmt.Println()

	num := 6
	fmt.Println(num)
	fmt.Println(&num)
	double(&num)
	fmt.Println(num)

	var myFloatPointer *float64 = createPointer()
	fmt.Println(*myFloatPointer)

	myBool := true
	printBoolPointer(&myBool)
}
