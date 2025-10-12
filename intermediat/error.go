package intermadiat

import (
	"errors"
	"fmt"
)

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	// Implementation of square root calculation (omitted for brevity)
	return 1, nil // Placeholder return value
}

func main() {

	result, err := sqrt(-16)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

}
