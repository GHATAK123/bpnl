package main

import (
	"errors"
	"fmt"
)

func AddProductToInventory(product Product) {
	inventory[product.ID] = product
}

func UpdateInventory(productID string, quantity int) error {
	product, exists := inventory[productID]
	if !exists {
		return errors.New("product not found")
	}
	if product.Quantity < quantity {
		return errors.New("insufficient quantity")
	}
	product.Quantity -= quantity
	inventory[productID] = product
	return nil
}

func ViewInventory() {
	for _, product := range inventory {
		fmt.Printf("ID: %s, Name: %s, Quantity: %d, Price: %.2f\n", product.ID, product.Name, product.Quantity, product.Price)
	}
}
