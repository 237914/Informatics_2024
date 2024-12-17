package Lab7

import (
	"fmt"
)

func GetFinalPrice(products []Product) float64 {
	var totalCost float64
	for _, product := range products {
		totalCost += product.getPrice()
	}
	return totalCost
}
func RunLab7() {
	product1 := &Device{Name: "Ноутбук", Price: 30, Brand: "Honor", Model: "1"}
	if err := product1.applyDiscount(25); err != nil {
		fmt.Println(err)
		return
	}
	product2 := &Cloth{Name: "Свитер", Price: 3.52, Size: "L", Color: "Красный"}
	if err := product2.applyDiscount(10); err != nil {
		fmt.Println(err)
		return
	}
	product3 := &Food{Name: "Арбуз", Price: 1.45, Weight: 1}
	if err := product3.applyDiscount(15); err != nil {
		fmt.Println(err)
		return
	}
	products := []Product{product1, product2, product3}
	for _, product := range products {
		fmt.Println(product.getProductInfo())
	}
	total := GetFinalPrice(products)
	fmt.Printf("Общая стоимость: %.2f\n", total)
}
