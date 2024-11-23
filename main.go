package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Welcome to the Go Calculator")
	fmt.Println("Enter the first number:")
	fmt.Scanln(&num1)

	fmt.Println("Enter an operator (+, -, *, /):")
	fmt.Scanln(&operator)

	fmt.Println("Enter the second number:")
	fmt.Scanln(&num2)

	switch operator {
	case "+":
		fmt.Printf("Result: %.2f\n", num1+num2)
	case "-":
		fmt.Printf("Result: %.2f\n", num1-num2)
	case "*":
		fmt.Printf("Result: %.2f\n", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("Result: %.2f\n", num1/num2)
		} else {
			fmt.Println("Error: Division by zero!")
		}
	default:
		fmt.Println("Invalid operator!")
	}
}
