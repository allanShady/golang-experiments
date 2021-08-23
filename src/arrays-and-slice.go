package main

import (
	"fmt"
	"sort"
)

func main() {

	//ar ages [3]int  = [3]int {12, 21, 23}

	var ages [3]int = [3]int{12, 21, 23}
	players := [4]string{"Messi", "Cristiano", "m'bape", "Neymar"}

	fmt.Println(ages, len(ages))
	fmt.Println(players, len(players))

	//slices
	var scores = []int{3, 5, 9}
	scores = append(scores, 23)

	fmt.Println(scores, len(scores))

	// ranges
	range_one := players[1:3]  // [m'bape Neymar]
	range_two := players[2:]   //[m'bape Neymar]
	range_three := players[:3] //[Messi Cristiano m'bape]

	fmt.Println(range_one, range_two, range_three)

	// sort package usage
	numbers := []int{4, 1, 30, 0, 90}
	sort.Ints(numbers)
	fmt.Println(numbers)

	zero_index := sort.SearchInts(numbers, 0)
	fmt.Println(zero_index)

	names := []string{"Salman", "Camilo", "Kelvin", "MJ"}
	sort.Strings(names)
	fmt.Println(names)

}
