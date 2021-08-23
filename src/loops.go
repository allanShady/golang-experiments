package main

import "fmt"

func main() {

	players := [4]string{"Messi", "Cristiano", "m'bape", "Neymar"}

	for position := 0; position < len(players); position++ {
		fmt.Println(players[position])
	}

	
}
