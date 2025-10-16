package intermediat

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("outout.txt")
	if err != nil {
		fmt.Println("Error greating filr: ", err)
		return
	}
	defer file.Close()

	data := []byte("Hello World\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
	}

	fmt.Println("Data has been written to file succesfully")
}
