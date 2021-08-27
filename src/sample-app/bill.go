package main

import (
	"fmt"
	"os"
)

type Bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func FactoryBill(name string) Bill {
	return Bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
}

func (bill *Bill) FormatOutput() string {
	output := "Bills break down: \n"
	var total_bills float64 = 0

	for key, value := range bill.items {
		output += fmt.Sprintf("%-25v ...$%0.2f\n", key+":", value)
		total_bills += value
	}

	output += fmt.Sprintf("%-25v ...$%v\n", "tip:", bill.tip)

	output += fmt.Sprintf("%-25v ...$%0.2f\n", "total:", total_bills+bill.tip)

	return output
}

// update tip
func (bill *Bill) UpdateTip(tip float64) {
	bill.tip = tip
}

// add item
func (bill *Bill) AddItem(name string, price float64) {
	bill.items[name] = price
}

func (bill *Bill) Save() {
	data := []byte(bill.FormatOutput())
	err := os.WriteFile("src/sample-app/bills/"+bill.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("Bill saved successful - ", bill.name)
}
