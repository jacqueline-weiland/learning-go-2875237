package main

import (
	"fmt"
)

const aConst int = 64

func main() {

	var aString string = "This is Go!"

	fmt.Println(aString)
	fmt.Printf("The variable's type is %T\n", aString)
	// needed to add a line feed so new line is made and 42 isn't just concatenated

	var anInteger int = 42
	fmt.Println(anInteger)

	var defaultInt int
	fmt.Println(defaultInt)

	var anotherString = "This is another string"
	fmt.Println(anotherString)
	fmt.Printf("This variables type is %T\n", anotherString)

	var anotherInt = 53
	fmt.Println(anotherInt)
	fmt.Printf("This variables type is %T\n", anotherInt)

	myString := "This is also a String"
	// := only works for variables defined in functions
	fmt.Println(myString)
	fmt.Printf("The variable's type is %T\n", myString)

	fmt.Println(aConst)
	fmt.Printf("The variable's type is %T\n", aConst)

}
