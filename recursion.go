package main

import "fmt"

func main() {
	fmt.Println("Factorial of 5:", factorial(5))
	fmt.Println("Fibonacci of 7:", fibonacci(7))

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
