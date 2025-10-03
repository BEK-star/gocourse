package intermediat

import "fmt"

func main() {
	//*a := adder()
	//fmt.Println(a()) // 1
	//fmt.Println(a()) // 2
	//fmt.Println(a()) // 3

	//b := adder()
	//fmt.Println(b()) // 1
	//fmt.Println(b()) // 2
	//fmt.Println(b())

	substract := func() func(int) int {
		total := 99
		return func(x int) int {
			total -= x
			return total
		}
	}

	s := substract()
	fmt.Println(s(1))
	fmt.Println(s(2))
	fmt.Println(s(3))
}

func adder() func() int {
	sum := 0
	fmt.Println("adder called, sum is", sum)
	return func() int {
		sum++
		fmt.Println("inner func called, 1 added to sum")
		return sum
	}
}
