package main

import "fmt"

func Greeting(name string) {
	fmt.Println("Hello, ", name)
}

func GreetingPlayers(players []string, f func(string)) {
	for _, player := range players {
		f(player)
	}
}
func main() {

	Greeting("Allan")
	GreetingPlayers([]string {"Ronaldo", "Messi", "Nani", "Levandosk"}, Greeting)

}
