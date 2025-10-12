package intermediat

import (
	"errors"
	"fmt"
)

type customError struct {
	Code    int
	Message string
	er      error
}

func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s, %v", e.Code, e.Message, e.er)
}
func doSomething() error {
	err := doSomethingElse()
	if nil != err {
		return &customError{Code: 404, Message: "Resource not found", er: err}
	}
	return nil
}

func doSomethingElse() error {
	return errors.New("unidentified error occurred")
}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Operation successful")

}
