package main

import "fmt"

func main() {

	players_ages := map[string]float32{
		"Ronaldo": 24,
		"Messi":   44,
		"Neymar":  22,
	}

	fmt.Println(players_ages)
	fmt.Println(players_ages["Neymar"])

	// Loop array
	for key, age := range players_ages {
		fmt.Println(key, "-", age)
	}
}
