package intermediat

import (
	"fmt"
)

func main() {
	str := "Hello, 世界"
	for i, r := range str {
		fmt.Printf("Character %d: %c (Unicode: %U)\n", i, r, r)
	}
}
