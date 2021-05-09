package main

import (
	"fmt"
	"github.com/TimLeary/headFirstGo/calendar"
	"log"
)

func main()  {
	date := calendar.Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Year())

	err = date.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Month())

	err = date.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Day())
}
