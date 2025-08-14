package basics

import "fmt"

func main() {

	/*fruit := "apple"
	switch fruit {
	case "apple":
		fmt.Println("it is an apple")
	case "banana":
		fmt.Println("It is a banana")
	default:
		fmt.Println("Unknown fruit")

	}

	day := "Monday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday":
		fmt.Println("It is a weekday")
	case "Sunday":
		fmt.Println("It is a weekend")
	default:
		fmt.Println("Unknown day")
	}

	/*number := 21
	switch {
	case number < 10:
		fmt.Println("Number is less than 10")
	case number >= 10 && number < 20:
		fmt.Println("Number is between 10 and 20")
	default:
		fmt.Println("Number is more than 20")
	}*/

	/*num := 2
	switch {
	case num > 1:
		fmt.Println("Greater than 1")
		fallthrough
	case num == 2:
		fmt.Println("Num is equal 2")
	default:
		fmt.Println("I do not know")
	}*/
	checkTiper(24)
	checkTiper("hello")
	checkTiper(true)

}

func checkTiper(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It is Integer")
	case string:
		fmt.Println("It is String")
	default:
		fmt.Println("Unknown")
	}
}
