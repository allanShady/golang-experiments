package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CreateBill() Bill {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Create a new Bill name: ")

	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	bill := FactoryBill(name)
	fmt.Println("Created the bill - ", bill)

	return bill
}

func main() {
	bill := CreateBill()

	fmt.Println(bill)
}
