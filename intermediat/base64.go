package intermediat

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("Hello, base64")

	transform := base64.StdEncoding.EncodeToString(data)
	fmt.Println("Encoded: ", transform)

	untransform, err := base64.StdEncoding.DecodeString(transform)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Decoded: ", string(untransform))
}
