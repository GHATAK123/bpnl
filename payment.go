package main

import (
	"errors"
	"fmt"
)

func ClearDues(userID string, amount float64) error {
	user, exists := users[userID]
	if !exists {
		return errors.New("user not found")
	}
	user.Credit += amount
	users[userID] = user
	fmt.Printf("Dues cleared for user %s. New credit balance: %.2f\n", user.Name, user.Credit)
	return nil
}
