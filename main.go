package main

import (
	"fmt"
	"unit_test/stringutil"
	"unit_test/calculator"
)

func main() {

	additionOp := calculator.GetOperation("addition")
	result := additionOp.Operate(9, 4)
	fmt.Println("9 + 4 =", result)

	multiplicationOp := calculator.GetOperation("multiplication")
	result = multiplicationOp.Operate(9, 4)
	fmt.Println("9 * 4 =", result)

	fmt.Println(stringutil.Reverse("alberd"))
	fmt.Println(stringutil.ToUpper("alberd"))
	fmt.Println(stringutil.IsEmpty("alberd"))
	fmt.Println(stringutil.FirstRune("alberd"))
	fmt.Println(stringutil.Concat("alberd"))
}