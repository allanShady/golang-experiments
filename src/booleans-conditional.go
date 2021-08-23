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

	//continue and break
	players := []string{"Messi", "Cristiano", "m'bape", "Neymar", "Tester"}

	for position, player := range players {
		if position == 1 {
			fmt.Printf("continue at position: %v and the value was: %s\n", position, player)
		} else if position > 2 {
			fmt.Printf("break at position %v\n", position)
			break
		}
	}
}
