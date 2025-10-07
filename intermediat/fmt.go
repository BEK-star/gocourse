package intermediat

import "fmt"

func main() {
	//*fmt.Print("Hello, World!1")
	//*fmt.Print("Hello, World!2")
	//*fmt.Print("Hello, World!3")

	//*fmt.Println("Hello, World4")
	//*fmt.Println("Hello, World!5")
	//*fmt.Println("Hello, World!6")

	//*name := "John"
	//*age := 30
	//*fmt.Printf("Name : %s, Age : %d\n", name, age)
	//*fmt.Printf("Name : %s, Age : %d\n", "Alice", 25)
	//*fmt.Printf("Binary : %b, Hex : %x\n", age, age)

	//*s := fmt.Sprint("hello Miz", "ooxas", 123, 24)
	//*fmt.Print(s)
	//*fmt.Print(s)

	//*c := fmt.Sprintln("hello Miz", "ooxas", 123, 24)
	//*fmt.Print(c)
	//*fmt.Print(c)

	//*d := fmt.Sprintf("Name : %s, Age : %d", name, age)
	//*fmt.Print(d)
	//*fmt.Println(d)

	//*var name string
	//*var age int

	//*fmt.Print("Enter your name: ")
	//*fmt.Scanln(&name, &age)

	//*fmt.Printf("Name : %s, Age : %d\n", name, age)

	//* handling error
	age := 19
	err := checkErr(age)
	if err != nil {
		fmt.Println("Error;", err)
	}

}

func checkErr(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d in too young", age)
	}
	return nil
}
