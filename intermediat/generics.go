package intermedat

import "fmt"

//func swap[T any](a, b T) (T, T) {
//	return b, a
//}

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(element T) {
	s.data = append(s.data, element)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	info := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return info, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) PrintAll() {
	if len(s.data) == 0 {
		fmt.Println("stack is empty")
		return
	}
	fmt.Print("stack elements: ")
	for _, v := range s.data {
		fmt.Print(v)
	}
	fmt.Println()
}

func main() {
	//a, b := 1, 2
	//a, b = swap(a, b)
	//fmt.Println(a, b)

	//c, d := "hello", "world"
	//c, d = swap(c, d)
	//fmt.Println(c, d)

	var intStack Stack[int]
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	intStack.PrintAll()

	intStack.Pop()
	intStack.PrintAll()
	intStack.Pop()
	intStack.Pop()
	intStack.PrintAll()
	fmt.Println("is empty:", intStack.IsEmpty())

	stringStack := Stack[string]{}
	stringStack.Push("hello ")
	stringStack.Push("world")
	stringStack.PrintAll()
	fmt.Println(stringStack.Pop())
	stringStack.PrintAll()
	stringStack.Push("Mike")
	stringStack.PrintAll()

}
