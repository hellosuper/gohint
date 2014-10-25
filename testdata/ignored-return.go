// Test that return values are processed

// Package pkg does something.
package pkg

import "errors"

type megaErr struct {
	error
}
func (i megaErr) Error() string {
	return "I am THE error"
}

func returnOne() int {
	return 0
}

func returnTwo() (int, string) {
	return 0, "something"
}

func returnErrOne() error {
	return errors.New("err")
}

func returnErrTwo() (asd int, err error) {
	return 0, errors.New("err")
}

// TODO: implement deep check for returned values that implement error
func returnErrStruct() (string, megaErr) {
	return "", megaErr{}
}


func main() {
	returnOne() // MATCH /result of 'returnOne' should not be silently ignored/
	returnTwo() // MATCH /result of 'returnTwo' should not be silently ignored/

	returnErrOne() // MATCH /function 'returnErrOne' returns an error, it should not be silently ignored/
	a, _ := returnErrTwo() // MATCH /function 'returnErrTwo' returns an error, generally it should not be intentionally ignored/

	if a, _ := returnErrTwo(); a > 0 { // MATCH /function 'returnErrTwo' returns an error, generally it should not be intentionally ignored/
		doSomethingMan()
	}

	//returnErrThree() // TODO: see above

}
