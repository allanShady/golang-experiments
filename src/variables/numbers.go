package main

import "fmt"

func main() {

	// integers
	var weight = 78
	var age int = 12
	var birth_day = 1970
	hight := 1.76

	fmt.Println("Greeting!", weight, age, ",", birth_day, "hight is", hight)

	// bits and memory
	var counter int8 = 12
	var negative_counter int8 = -128
	var unsigned_counter uint = 512

	fmt.Println(counter, negative_counter, unsigned_counter)

	// floats
	var rate_limit float32 = 234.12
	average := 772227.8888823

	fmt.Println("This value comes from a float var of 32 bits:", rate_limit)
	fmt.Println("This value comes from a float var of 64 bits:", average)

}
