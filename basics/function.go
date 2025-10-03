package basics

import "fmt"

func main() {

	greed := func() {
		fmt.Println("Hello Anonumus function")
	}

	fmt.Println(add(1, 2))
	greed()
}

func add(a, b int) int {
	return a + b

}
