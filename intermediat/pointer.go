package intermadiat

import "fmt"

func main() {
	x := 10

	modifyValue(&x)
	fmt.Println(x)
	modifyValue2(&x)
	fmt.Println(x)

}

func modifyValue(a *int) {
	*a++
}

func modifyValue2(a *int) {
	*a++
}
