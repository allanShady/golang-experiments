package main

import (
	"fmt"
	"math"
)

func Greeting(name string) {
	fmt.Println("Hello, ", name)
}

func GreetingPlayers(players []string, f func(string)) {
	for _, player := range players {
		f(player)
	}
}

func CircleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func main() {

	Greeting("Allan")
	GreetingPlayers([]string{"Ronaldo", "Messi", "Nani", "Levandosk"}, Greeting)

	circle_area := CircleArea(9)
	fmt.Printf("The final area is: %0.2f\n", circle_area)
}
