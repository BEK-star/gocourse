package intermediat

import "fmt"

type Rectangle struct {
	width  float64
	length float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r *Rectangle) Scale(factor float64) {
	r.length = r.length * factor
	r.width = r.width * factor
}

func main() {

	rect := Rectangle{
		width:  9,
		length: 5,
	}

	ar := rect.Area()
	fmt.Println("Area before scaling:", ar)

	rect.Scale(2)
	ar = rect.Area()

	fmt.Println("Area after scaling:", ar)

}
