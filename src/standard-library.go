package main

import (
	"fmt"
	"strings"
)

func main() {
	greeting := "Hi there!"

	fmt.Println(strings.Contains(greeting, "Hi")) //true
	fmt.Println(strings.ReplaceAll(greeting, "Hi", "Hello")) //Hello there!
	fmt.Println(strings.ToUpper(greeting)) //HI THERE!
	fmt.Println(strings.Index(greeting, "!")) //8
	fmt.Println(strings.Split(greeting, " ")) //[Hi there!]

}
