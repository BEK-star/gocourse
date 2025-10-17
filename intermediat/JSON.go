package inermediat

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	Age       int    `json:"age"`
}

func main() {
	person := Person{FirstName: "John", Age: 30}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error, making json format", err)
		return
	}
	fmt.Println(string(jsonData))
}
