package main

import "fmt"

func Greeting(name string) {
	fmt.Println("name memory ref is:", &name)

	// set poiter
	name_pointer := &name
	// get the pointer value
	fmt.Println(*name_pointer)
}

func UpdateName(name *string) {
	*name = "John Doe"
}

func main() {
	name := "Allan Camilo"

	Greeting(name)
	UpdateName(&name)
	fmt.Println("Now the name is not 'Allan Camilo' anymore and it is: ", name)
}
