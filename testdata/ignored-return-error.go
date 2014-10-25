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

func returnOne() error {
	return errors.New("err")
}

func returnTwo() (asd int, err error) {
	return 0, errors.New("err")
}

// TODO: implement deep check for returned values that implement error
func returnThree() (string, megaErr) {
	return "", megaErr{}
}

func main() {
	returnOne() // MATCH /function 'returnOne' returns an error, it should not be silently ignored/
	a, _ := returnTwo() // MATCH /function 'returnTwo' returns an error, generally it should not be intentionally ignored/

	if a, _ := returnTwo(); a > 0 { // MATCH /function 'returnTwo' returns an error, generally it should not be intentionally ignored/
		doSomethingMan()
    }

	//returnThree() // TODO: see above
}
