package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	//contact   contactInfo
	// We can also declare the contactInfo property this way. It means there's a property called contactInfo of type contactInfo
	contactInfo
}

func main() {
	// There are multiple ways to initialize a struct

	// a)
	// No need to specify the properties, Go assumes the order based on the definition of the struct
	//alex := person{"Alex", "Anderson"}

	// b)
	//alex := person{firstName: "Alex", lastName: "Anderson"}

	// c)
	// var alex person
	// At this point, the default values are in place since we haven't initialize the properties
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)
	// fmt.Println()

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// &variable = Give me the memory address of the value this variable is pointing at
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()
}

// Notice how we changed the receiver type here to be *person
// This is a type description - it means we are working with a pointer of type person
func (pointerToPerson *person) updateName(newFirstName string) {
	// *pointer Give me the value this memory address is pointing at
	// This is an operator - it means we want to manipulate the value the pointer is referencing
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
