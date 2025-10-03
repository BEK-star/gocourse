package basics

import (
	"errors"
	"fmt"
)

func main() {

	q, r := divide(3, 3)
	fmt.Printf("Quotion : %d   Reminder : %d\n", r, q)

	result, err := compare(3, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

}

func divide(a, b int) (int, int) {
	quotion := a / b
	reminder := a % b
	return quotion, reminder

}

func compare(a, b int) (string, error) {
	if a > b {
		return "A is greater", nil
	} else if a < b {
		return "B is greater", nil
	} else {
		return "", errors.New("Cannot compare equal values")
	}
}
