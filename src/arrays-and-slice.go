package main

import "fmt"

func main() {

	//ar ages [3]int  = [3]int {12, 21, 23}

	var ages [3]int = [3]int{12, 21, 23}
	players := [4]string{"Messi", "Cristiano", "m'bape", "Neymar"}

	fmt.Println(ages, len(ages))
	fmt.Println(players, len(players))
}
