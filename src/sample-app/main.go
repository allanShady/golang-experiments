package main

import "fmt"

func main() {
	sample_bill := FactoryBill("Allan's bill")
	fmt.Println(sample_bill.FormatOutput())
}
