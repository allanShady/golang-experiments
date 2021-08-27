package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func PromptOptions(bill Bill) {
	reader := bufio.NewReader(os.Stdin)

	chosenOption, _ := GetInput("Choose an option (a - add item, s - save item, t - add tip): ", reader)

	switch chosenOption {
	case "a":
		name, _ := GetInput("Item name: ", reader)
		price, _ := GetInput("Item price: ", reader)

		priceAsFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be number")
			PromptOptions(bill)
		}

		bill.AddItem(name, priceAsFloat)
		fmt.Println("Item added successful: ", name, price)

		PromptOptions(bill)
	case "t":
		tip, _ := GetInput("Enter tip amount ($): ", reader)

		tipAsFloat, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be number")
			PromptOptions(bill)
		}

		bill.UpdateTip(tipAsFloat)
		fmt.Println("Tip added successful: ", tipAsFloat)

		PromptOptions(bill)
		fmt.Println(tip)
	case "s":
		fmt.Println("The following bill was successful saved ", bill)
	default:
		fmt.Println("Oops! invlaid chosen option. Please try again.")
		PromptOptions(bill)
	}
}

func main() {
	bill := CreateBill()
	PromptOptions(bill)
}
