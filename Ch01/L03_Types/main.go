package main

import (
	"fmt"
	"reflect"
)

func main()  {
	fmt.Println('Җ')
	fmt.Println(true, reflect.TypeOf(true))
	fmt.Println(42, reflect.TypeOf(42))
	fmt.Println(3.1415, reflect.TypeOf(3.1415))
	fmt.Println("Hello Go!", reflect.TypeOf("Hello Go!"))
}
