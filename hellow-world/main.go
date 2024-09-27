package main

import "fmt"

const constVar int = 54

func main() {

	fmt.Println("Hellow Go !")

	// decalre varaible
	var string1 string = "Lovely Morning !"
	var int1 int = 10
	fmt.Println(string1)
	fmt.Println(int1)

	// Get type of varaible
	fmt.Printf("The variable's type is %T\n", string1)
	fmt.Printf("The variable's type is %T\n", constVar)

	// no need to declare the type
	var string2 = "Good Evening !"
	var int2 = 20
	string3 := "Good Night !"

	fmt.Println(string2, int2, string3)
}
