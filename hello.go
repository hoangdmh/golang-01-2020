package main

import (
	"fmt"
)

var i int = 2

func add() int {
	sum := 0
	for i := 0; i < 4; i++ {
		sum += i
	}
	return sum
}

func switchFun(i int) int {
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("zero")
	}
	return i
}

func main() {
	fmt.Println("result is", add())
	fmt.Println("i", i)
	fmt.Println("switch is", switchFun(1))
	fmt.Printf("Type: %T Value: %v\n", i, i)
}
