package main

import (
	"fmt"
)

type employee struct {
	name    Name
	age     int
	city    string
	country string
}

type Name struct {
	fName string
	lName string
}

func main() {
	emp := employee{
		Name{
			"Ab",
			"cd",
		},
		22,
		"mumbai",
		"India",
	}
	fmt.Println("First Name:", emp.name.fName)
	fmt.Println("Last Name:", emp.name.lName)
	fmt.Println("Age:", emp.age)
	fmt.Println("City:", emp.city)
	fmt.Println("Country:", emp.country)
}
