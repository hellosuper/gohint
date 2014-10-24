// Test that return values are processed

// Package pkg does something.
package pkg

func returnOne() int {
	return 0
}

func returnTwo() (int, string) {
	return 0, "something"
}

func main() {
	returnOne() // MATCH /ignored function return: int/
	returnTwo() // MATCH /ignored function return: int, string/
}
