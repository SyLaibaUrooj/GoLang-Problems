package main
import (
	"fmt"
"strconv"
)

// Change these boolean values to control whether you see 
// the expected answer and/or hints.
const showExpectedResult = false;
const showHints = false;

// calculate() returns the sum of the two parameters
func calculate(value1 string, value2 string) float64 {
    // Your code goes here.
	var sum float64 = 0
	// Convert the first string to a float64
	valueFloat1, err := strconv.ParseFloat(value1, 64)

	// Convert the second string to a float64
	valueFloat2, err := strconv.ParseFloat(value2, 64)

	// error 
	if err != nil {
		fmt.Println("Error converting string to float64:", err)
		return -1
	}

	// Calculate and return the result
	sum = valueFloat1 + valueFloat2


	return sum
}
