package lab7

import (
	"fmt"
)

type food struct {
	Name   string
	Price  float64
	Weight float64
}

func (f *food) applyDiscount(discount float64) error {
	if err := validateDiscount(discount); err != nil {
		return err
	}
	f.Price = f.Price * (1 - discount/100)
	return nil
}
func (f *food) getPrice() float64 {
	return f.Price
}
func (f *food) getProductInfo() string {
	if f.Name == "" {
		return ("Не все поля заполнены")
	}
	return fmt.Sprintf("Название: %s, Вес: %.2f кг, Цена: %.2f", f.Name, f.Weight, f.Price)
}
