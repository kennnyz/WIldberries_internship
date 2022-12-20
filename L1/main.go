package main

import "fmt"

type ProductList []Product

func (products *ProductList) calcTotalCategory() map[string]float64 {
	totals := make(map[string]float64)
	for _, p := range *products {
		totals[p.category] = totals[p.category] + p.price
	}
	return totals
}

func getProducts() []Product {
	return []Product{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
}

func main() {
	products := ProductList(getProducts())
	for category, total := range products.calcTotalCategory() {
		fmt.Println("Category: ", category, "Total: ", total)
	}
}
