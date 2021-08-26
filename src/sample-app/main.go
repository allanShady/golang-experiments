package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, error := reader.ReadString('\n')

	return strings.TrimSpace(input), error

}

func CreateBill() Bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := GetInput("Create a new Bill name: ", reader)

	bill := FactoryBill(name)
	fmt.Println("Created the bill - ", bill)

	return bill
}

func main() {
	bill := CreateBill()

	fmt.Println(bill)
}
