package main

import "fmt"

func main() {
	var pointerVariable *int     // var is nill
	fmt.Println(pointerVariable) // prints <nil>

	aInt := 51
	pointerVariable = &aInt // pointerVar is now pointing to aInt

	fmt.Println(*pointerVariable) // prints 51

}
