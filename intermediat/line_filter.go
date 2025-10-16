package intermediat

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("outout.txt")
	if err != nil {
		fmt.Println("Error, opening file: ", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	keyword := "important"

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			fmt.Println("Filtered file:", line)
		}

	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Errror scanning file: ", file)
		return
	}

}
