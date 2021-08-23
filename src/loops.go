package main

import "fmt"

func main() {

	players := [4]string{"Messi", "Cristiano", "m'bape", "Neymar"}

	for position := 0; position < len(players); position++ {
		fmt.Println(players[position])
	}

	//get index and a copy of the value
	for position, player := range players {
		fmt.Printf("The player %v is %s \n", position+1, player)
	}
}
