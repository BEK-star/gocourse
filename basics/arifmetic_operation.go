package basics

import "fmt"

func main() {
	var a, b int = 10, 5
	var result int

	fmt.Println("Numbers:", a, "|", b)

	fmt.Println("_____________")

	result = a + b
	fmt.Println("Addition =", result)

	result = a - b
	fmt.Println("Subtraction =", result)

	result = a * b
	fmt.Println("Multiplication =", result)

	result = a / b
	fmt.Println("Division =", result)
}
