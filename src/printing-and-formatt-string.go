package main

import "fmt"

func main() {

	default_year := 1970
	os_kernel := "Unix"

	// printf (formatt strings) %_ = formatt specifier
	fmt.Printf("OS kernel: %v default year: %v \n", os_kernel, default_year)
	fmt.Printf("OS kernel: %q default year: %v \n", os_kernel, default_year)

	fmt.Printf("default year is type of: %T \n", default_year)

	fmt.Printf("formatting float value: %f \n", 24.8924)
	fmt.Printf("formatting  float value: %0.2f \n", 24.8924)

	var formatted_value = fmt.Sprintf("OS kernel: %v default year: %v \n", os_kernel, default_year)
	fmt.Println(formatted_value)
}
