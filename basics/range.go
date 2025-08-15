package basics

import "fmt"

func main() {
	message := "Hello world"

	for i, v := range message {
		fmt.Printf("Index %d   Value %c\n", i, v)
	}
}
