package main

import (
	"fmt"
	"github.com/TimLeary/headFirstGo/datafile"
	"log"
)

func main()  {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0
	for _, number := range numbers{
		sum += number
	}

	fmt.Printf("Average: %0.2f\n", sum / float64(len(numbers)))
}
