package main

import "fmt"

func main() {

	AddProductToInventory(Product{ID: "p1", Name: "Books", Quantity: 10, Price: 1000.00})
	AddProductToInventory(Product{ID: "p2", Name: "Pen", Quantity: 20, Price: 500.00})

	users["user1"] = User{ID: "u1", Name: "Prakash", Credit: 2000.00}

	err := PlaceOrder("user1", "p1", 1, "Prepaid")
	if err != nil {
		fmt.Println(err)
	}

	err = PlaceOrder("user1", "p2", 1, "BNPL")
	if err != nil {
		fmt.Println(err)
	}

	ViewOrderHistory("user1")

	err = ClearDues("user1", 500)
	if err != nil {
		fmt.Println(err)
	}

	ViewInventory()

	// Check for defaulted orders
	CheckForDefaultedOrders("user1")
}

type Product struct {
	ID       string
	Name     string
	Quantity int
	Price    float64
}

type Order struct {
	ID          string
	UserID      string
	ProductID   string
	Quantity    int
	Total       float64
	PaymentMode string
	Status      string
}

// Struct for User with credit and order history
type User struct {
	ID             string
	Name           string
	Credit         float64
	Orders         []Order
	Blacklist      bool
	DefaultedCount int
}

// Inventory map
var inventory = make(map[string]Product)

// User map
var users = make(map[string]User)

// Order map
var orders = make(map[string]Order)
