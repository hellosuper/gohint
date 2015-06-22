// Test that return values has no names

// Package foo ...
package foo

func f1() (x int) { // MATCH /return value.*should not be named/
	return 0
}

func f2() (x, y int) { // MATCH /return value.*should not be named/
	return 0, 0
}

func f3() (x int, y int) { // MATCH /return value.*should not be named/
	return 0, 0
}

func f4() (int, y int) { // MATCH /return value.*should not be named/
	return 0, 0
}

type ret struct{}

func (r ret) f5() (x int) { // MATCH /return value.*should not be named/
	return 0
}

func (r ret) f6() (x, y int) { // MATCH /return value.*should not be named/
	return 0, 0
}

func (r ret) f7() (x int, y int) { // MATCH /return value.*should not be named/
	return 0, 0
}

func (r ret) f8() (int, y int) { // MATCH /return value.*should not be named/
	return 0, 0
}

func (r ret) f9() (int int) { // MATCH /return value.*should not be named/
	return 0, 0
}

func (r ret) f10() (int, y int) { // MATCH /return value.*should not be named/
	return 0, 0
}
