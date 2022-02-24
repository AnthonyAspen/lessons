package main

import (
	"log"
)

type Interface interface {
	Square(int) (int, error)
	Sum(int) (int, error)
}

type Math struct {
	integer int
}

func (m *Math) Square(i int) (int, error) {
	return i * i, nil
}

func (m *Math) Sum(i int) (int, error) {
	return i + i, nil
}

// create mock struct to pass it to testMath
type Mock struct {
	integer int
}

func (m *Mock) Square(i int) (int, error) {
	return 0, nil
}

func (m *Mock) Sum(i int) (int, error) {
	return 0, nil
}

func main() {
	trueStruct := &Math{integer: 10}
	mockStruct := &Mock{integer: 10}

	log.Printf("TrueValue:%v,MockValue:%v", testMath(trueStruct, 10), testMath(mockStruct, 10))

}

func testMath(Interface Interface, i int) int {
	sq, err := Interface.Square(i)
	if err != nil {
		return 0
	}
	su, err := Interface.Sum(i)
	if err != nil {
		return 0
	}
	return sq + su

}
