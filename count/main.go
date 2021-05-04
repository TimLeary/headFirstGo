package main

import (
	"fmt"
	"github.com/TimLeary/headFirstGo/datafile"
	"log"
)

func main()  {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}
