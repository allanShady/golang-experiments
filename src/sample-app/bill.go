package main

import "fmt"

type Bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func FactoryBill(name string) Bill {
	return Bill{
		name:  name,
		items: map[string]float64{"coffee": 5.087, "xima": 21.879},
		tip:   0,
	}
}

func (bill Bill) FormatOutput() string {
	output := "Bills break down: \n"
	var total_bills float64 = 0

	for key, value := range bill.items {
		output += fmt.Sprintf("%-25v ...$%0.2f\n", key+":", value)
		total_bills += value
	}

	output += fmt.Sprintf("%-25v ...$%0.2f\n", "total:", total_bills)

	return output
}
