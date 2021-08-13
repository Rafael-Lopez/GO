package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// There are multiple ways to initialize a struct

	// a)
	// No need to specify the properties, Go assumes the order based on the definition of the struct
	//alex := person{"Alex", "Anderson"}

	// b)
	//alex := person{firstName: "Alex", lastName: "Anderson"}

	// c)
	var alex person
	// At this point, the default values are in place since we haven't initialize the properties
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
	fmt.Println()

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)
}
