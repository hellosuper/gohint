// Test that return values are processed

// Package pkg does something.
package pkg

import "errors"

func returnOne() error {
	return errors.New("err")
}

func returnTwo() (int, error) {
	return 0, errors.New("err")
}

func main() {
	returnOne() // MATCH /unprocessed returned error/
	a, _ := returnTwo() // MATCH /unprocessed returned error: int, error/
}
