package main

import "fmt"

func main() {
	bill := FactoryBill("Allan's bill")

	bill.AddItem("coffee puding", 7.321)

	bill.UpdateTip(40)
	fmt.Println(bill.FormatOutput())
}
