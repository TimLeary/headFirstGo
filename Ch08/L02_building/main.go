package main

import (
	"fmt"
	"github.com/TimLeary/headFirstGo/magazine"
)

func main()  {
	var s magazine.Subsciber
	s.Rate = 4.99
	fmt.Println(s.Rate)
}
