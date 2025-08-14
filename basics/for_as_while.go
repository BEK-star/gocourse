package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}*/

	sourse := rand.NewSource(time.Now().UnixNano())
	random := rand.New(sourse)

	target := random.Intn(100) + 1

	var guess int
	for {
		fmt.Println("Enter number :")
		fmt.Scan(&guess)

		if target == guess {
			fmt.Println("you Win!!!")
			break
		} else if target < guess {
			fmt.Println("High number")
		} else {
			fmt.Println("low number")
		}
	}
}
