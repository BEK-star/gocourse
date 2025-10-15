package intermediat

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// reader := bufio.NewReader(strings.NewReader("Hello world, My name is John\n"))

	// data := make([]byte, 20)

	// n, err := reader.Read(data)
	// if err != nil {
	// 	fmt.Println("Error, reading", err)
	// 	return
	// }
	// fmt.Printf("Read %d bytes: %s \n", n, data[:n])

	// line, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Error, reading: ", err)
	// 	return
	// }
	// fmt.Printf("Remaining words:%s", line)

	writer := bufio.NewWriter(os.Stdout)

	data := []byte("Hello bufer !!!\n")

	line, err := writer.Write(data)
	if err != nil {
		fmt.Println("Error, write is:", err)
	}
	fmt.Printf("Byte %d ", line)

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error, flushing fell:", err)
	}
}
