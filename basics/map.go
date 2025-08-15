package basics

import "fmt"

func main() {
	//First way to declare the map with make
	myMap := make(map[string]string)
	myMap["Russia"] = "Moscow"
	myMap["German"] = "Berlin"

	fmt.Println("Capital of ", myMap["Russia"])

	//Second wsy to declare the map
	map2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(map2)

	for k, v := range map2 {
		fmt.Println(k, v)
	}
}
