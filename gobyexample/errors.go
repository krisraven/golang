package main
// in Go its idomatic to communicate errors via an explicit separate retrun value. This approach makes it easy to see which functions return errors 
import (
	"errors"
	"fmt"
)

// by default errors are the last return value and have type error (a built in interface)
// errors.New constructs a basic errorvalue with the given error message
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	// a nil value in the error position indicates that the was no error
	return arg + 3, nil
}

// Custom Errors can be used
type argError struct {
	arg		int
	prob	string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}
// &argError syntax builds a new struct, supplying values for the 2 fiels arg and prob
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}
//  the two loops test out each of the error returning functions. The inline error check on the if line is a common idiom in Go
func main() {
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed: ", e)
		} else {
			fmt.Println("f2 worked: ", r)
		}
	}
// can programmatically use the data in a custom error by getting the error as an instance of the custom error type via type assertion
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}