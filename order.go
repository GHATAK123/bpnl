package main

import (
	"errors"
	"fmt"
)

func PlaceOrder(userID, productID string, quantity int, paymentMode string) error {
	user, userExists := users[userID]
	if !userExists {
		return errors.New("user not found")
	}
	product, productExists := inventory[productID]
	if !productExists {
		return errors.New("product not found")
	}
	if product.Quantity < quantity {
		return errors.New("insufficient product quantity")
	}
	total := float64(quantity) * product.Price
	order := Order{
		ID:          fmt.Sprintf("%s-%s", userID, productID),
		UserID:      userID,
		ProductID:   productID,
		Quantity:    quantity,
		Total:       total,
		PaymentMode: paymentMode,
		Status:      "Placed",
	}

	if paymentMode == "Prepaid" {
		user.Orders = append(user.Orders, order)
		UpdateInventory(productID, quantity)
		orders[order.ID] = order
		fmt.Println("Order placed successfully with Prepaid mode.")
	} else if paymentMode == "BNPL" {
		if user.Blacklist {
			return errors.New("user is blacklisted, cannot use BNPL")
		}
		if user.Credit < total {
			return errors.New("insufficient credit")
		}
		user.Credit -= total
		user.Orders = append(user.Orders, order)
		UpdateInventory(productID, quantity)
		orders[order.ID] = order
		fmt.Println("Order placed successfully with BNPL.")
	} else {
		return errors.New("invalid payment mode")
	}
	users[userID] = user
	return nil
}

func ViewOrderHistory(userID string) {
	user, exists := users[userID]
	if !exists {
		fmt.Println("User not found")
		return
	}
	fmt.Printf("Order history for user %s:\n", user.Name)
	for _, order := range user.Orders {
		fmt.Printf("Order ID: %s, Product ID: %s, Quantity: %d, Total: %.2f, Payment Mode: %s, Status: %s\n",
			order.ID, order.ProductID, order.Quantity, order.Total, order.PaymentMode, order.Status)
	}
}
