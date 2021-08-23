package main

import "fmt"

func main() {
	age := 1

	if age > 18 {
		fmt.Println("You are a children")
	} else if age < 2 {
		fmt.Println("very little kid")
	} else {
		fmt.Println("Not found")
	}
}
