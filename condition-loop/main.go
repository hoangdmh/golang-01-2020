package main

import "fmt"

// condition statement
func conditionFunc(number int) {
	if number > 10 {
		fmt.Printf("%d lon hon", number)
	} else {
		fmt.Printf("%d be hon", number)
	}
}

//switch case
func switchCaseFunc(number int) {
	switch number {
	case 1:
		fmt.Printf("\n number is 1")
	case 2:
		fmt.Printf("\n number is 2")
	case 3, 4, 5, 6:
		fmt.Printf("\n number is %d", number)
	default:
		fmt.Printf("\n number is %d", number)
	}
}

//defer
func deferFunc() {
	defer fmt.Print(" word")
	fmt.Print("\n Hello")
}

// for loop
func forFunc() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("\n", sum)
}

func main() {
	fmt.Println("===========> Condition statement loop <===========")
	number := 6
	//return typeof a variable
	fmt.Printf("%T \n", number)

	// if else
	conditionFunc(number)

	//switch case
	switchCaseFunc(number)

	// defer => trì hoản thực thi 1 function cho đến khi các func xung quanh thực thi xong
	deferFunc()

	//for loop
	forFunc()

	fmt.Println("\n ===========> End Condition statement loop <===========")
}
