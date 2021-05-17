package main

import "fmt"

func calmDown() {
	fmt.Println(recover())

	// OR
	// p := recover()
	// err, ok := p.(error)
	// if ok {
	//	fmt.Println(err.Error())
	// }
}

func freakOut() {
	defer calmDown()
	panic("oh no, oh no ...")
	fmt.Println("I won't be run")

	// OR
	// defer calmDown()
	// err := fmt.Errorf("there's an error")
	// panic(err)
}

func main() {
	freakOut()
	fmt.Println("Exiting normal")
}
