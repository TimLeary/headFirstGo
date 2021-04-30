package main

import (
	"fmt"
	"github.com/TimLeary/headFirstGo/keyboard"
	"log"
)

func main() {
	fmt.Println("Enter a grade:")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60.0 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status)
}
