package main

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func FactoryBill(name string) bill {
	return bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
}
