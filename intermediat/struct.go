package intermediat

import "fmt"

type Person struct {
	Name       string
	SecondName string
	Age        int
	Address    Address
	PhoneNumber
}
type Address struct {
	City    string
	Country string
}

type PhoneNumber struct {
	CountryCode string
	Number      string
}

func main() {

	p := Person{
		Name:       "Alice",
		SecondName: "Smith",
		Age:        30,
		Address: Address{
			City:    "New York",
			Country: "USA",
		},
		PhoneNumber: PhoneNumber{
			CountryCode: "+1",
			Number:      "1234567890",
		},
	}

	fmt.Println(p.Name, p.Age)
	fmt.Println(p.fullname(), p.Age)
	p.addAge(2)
	fmt.Println(p.fullname(), p.Age, p.Address.City)
}

func (p Person) fullname() string {
	return p.Name + " " + p.SecondName
}

func (p *Person) addAge(Age int) {
	p.Age += Age
}
