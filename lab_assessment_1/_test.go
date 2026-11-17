//Build a Go inventory management module integrating all Go concepts from Sessions 1–4.
//Requirements:
//(1) define an Item struct with fields (e.g., ID, Name, Quantity, Price) and use a slice to store items;
//(2) implement add, update, and remove operations with input validation and proper error handling;
//(3) add a concurrent goroutine using channels to compute summary statistics (e.g., total value, item count); and
//(4) write at least three unit tests in a _test.go file for core functions

package main

import (
	"testing"
)

func TestAddPriceofallgivenitems(t *testing.T) {
	item := Item{ID: 1, Name: "Item1", Quantity: 10, Price: 5}
	err := item.AddPriceofallgivenitems(10)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if item.Price != 10 {
		t.Errorf("Expected price to be 10, but got %v", item.Price)
	}
}

func TestUpdateQuantity(t *testing.T) {
	item := Item{ID: 1, Name: "Item1", Quantity: 10, Price: 5}
	err := item.UpdateQuantity(15)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if item.Quantity != 15 {
		t.Errorf("Expected quantity to be 15, but got %v", item.Quantity)
	}
}

func TestRemoveItem(t *testing.T) {
	items := []Item{
		{ID: 1, Name: "Item1", Quantity: 10, Price: 5},
		{ID: 2, Name: "Item2", Quantity: 20, Price: 10},
	}
	item := Item{ID: 1, Name: "Item1", Quantity: 10, Price: 5}
	_, err := item.RemoveItem(items)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
