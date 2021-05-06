package main

import (
	"fmt"
	"github.com/TimLeary/headFirstGo/magazine"
)

func main()  {
	var subscriber magazine.Subscriber
	subscriber.Rate = 4.99
	fmt.Println(subscriber.Rate)

	var employee magazine.Employee
	employee.Name = "Joy Carr"
	employee.Salary = 60000
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)

	var address magazine.Address
	address.Street = "123 Oak St."
	address.City = "Omaha"
	address.PostalCode = "68111"

	subscriber.Address = address
	fmt.Println(subscriber.Address)

	employee.Street = "456 Elm St"
	employee.City = "Portland"
	employee.State = "OR"
	employee.PostalCode = "97222"
	fmt.Println(employee.Address)
}
