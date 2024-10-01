package main

import "fmt"

// Converts a slice of strings to a map object. Map will be Fruits with their prices. Prices will be divided such that they always add up to 100.
func convertToMap(items []string) map[string]float64 {

	// Create a map object
	result := make(map[string]float64)
	// Your code goes here
	var amt float64 = 100.0 / float64(len(items))
	for i := 0; i < len(items); i++ {
		result[items[i]] = float64(amt)
	}
	return result
}

func main() {
	slice := []string{"apple", "banana", "orange"}
	result := convertToMap(slice)
	fmt.Println(result)
}
