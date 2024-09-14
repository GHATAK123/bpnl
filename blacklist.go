package main

import "fmt"

func CheckForDefaultedOrders(userID string) {
	user, exists := users[userID]
	if !exists {
		fmt.Println("User not found")
		return
	}
	if user.DefaultedCount >= 3 {
		user.Blacklist = true
		users[userID] = user
		fmt.Printf("User %s has been blacklisted.\n", user.Name)
	} else {
		fmt.Println("User is not blacklisted.")
	}
}
