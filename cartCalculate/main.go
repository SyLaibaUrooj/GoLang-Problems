package main

import "fmt"

type cartItem struct {
	name     string
	price    float64
	quantity int
}

// calculateTotal() returns the total value of the shopping cart.
func calculateTotal(cart []cartItem) float64 {
	fmt.Println(cart)
	var totalPrice float64 = 0.0
	for i := 0; i < len(cart); i++ {
		for j := 0; j < cart[i].quantity; j++ {
			totalPrice += cart[i].price
		}
	}
	return totalPrice
}

func main() {
	var cart []cartItem
	var apples = cartItem{"apple", 2.99, 3}
	var oranges = cartItem{"orange", 1.50, 8}
	var bananas = cartItem{"banana", .49, 12}
	cart = append(cart, apples, oranges, bananas)
	result := calculateTotal(cart)
	fmt.Println("Total:", result)
}
