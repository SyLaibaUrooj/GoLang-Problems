// Write your answer here, and then test your code.

package main

import (
	"fmt"
	"log"
	"strconv"
)

// calculate() returns the result of the requested operation.
func calculate(input1 string, input2 string, operation string) float64 {
	var value1 = convertInputToValue(input1)
	fmt.Println("value1: ", value1)
	var value2 = convertInputToValue(input2)
	fmt.Println("value2: ", value2)

	switch operation {
	case "+":
		return addValues(value1, value2)
	case "-":
		return subtractValues(value1, value2)
	case "*":
		return multiplyValues(value1, value2)
	case "/":
		return divideValues(value1, value2)
	default:
		return 0
	}
}

func convertInputToValue(input string) float64 {
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatalf("Error converting input to float64: %v", err)
	}
	return value
}

func addValues(value1, value2 float64) float64 {
	fmt.Println("addValues: ", value1+value2)
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}

func main() {
	value1 := "10"
	value2 := "5.5"
	operation := "+"
	result := calculate(value1, value2, operation)
	println(result)
}
