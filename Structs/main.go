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
	alex := person{firstName: "Alex", lastName: "Anderson"}

	fmt.Println(alex)
}
