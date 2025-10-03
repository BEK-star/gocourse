package basics

import "fmt"

func main() {

	progress()
	fmt.Println("Program returned to normal execution")
}

func progress() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	fmt.Println("Starting the program")
	panic("Something went wrong!")
	fmt.Println("This line will not be executed")
}
