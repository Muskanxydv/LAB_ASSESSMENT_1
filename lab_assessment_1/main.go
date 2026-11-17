//Build a Go inventory management module integrating all Go concepts from Sessions 1–4.
//Requirements:
//(1) define an Item struct with fields (e.g., ID, Name, Quantity, Price) and use a slice to store items;
//(2) implement add, update, and remove operations with input validation and proper error handling;
//(3) add a concurrent goroutine using channels to compute summary statistics (e.g., total value, item count); and
//(4) write at least three unit tests in a _test.go file for core functions

package main

import "fmt"

type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    int
}

func (i *Item) TotalValue() int {
	return (i.Quantity) * i.Price
}

func (i *Item) Itemcount() int {
	return i.Quantity
}

func (i *Item) UpdateQuantity(newQuantity int) error {
	if newQuantity < 0 {
		return fmt.Errorf("Quantity of items cannot be negative!! Enter a positive value")
	}
	i.Quantity = newQuantity
	return nil
}

func (i *Item) RemoveItem(items []Item) ([]Item, error) {
	for index, item := range items {
		if item.ID == i.ID {
			return append(items[:index], items[index+1:]...), nil
		}
	}
	return items, fmt.Errorf("item cannot be found in the available items!! ENter a valid Item")
}

func (i *Item) AddPriceofallgivenitems(newPrice int) error {
	if newPrice <= 0 {
		return fmt.Errorf("Price of items cannot be zero or negative!!! Enter a positive value")
	}
	i.Price = newPrice
	return nil
}

func main() {
	item := Item{ID: 1, Name: "iPhone 17", Quantity: 3, Price: 130000}
	fmt.Println("Total Value:", item.TotalValue())
	fmt.Println("Item Count:", item.Itemcount())
	err := item.UpdateQuantity(5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Updated Quantity:", item.Quantity)
	}
}
